package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/cgiser/pushworker/common"
	"github.com/cgiser/pushworker/logging"
	"github.com/cgiser/pushworker/metrics"
	"github.com/cgiser/pushworker/task"
	jsoniter "github.com/json-iterator/go"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/urfave/cli"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"gopkg.in/go-playground/pool.v3"
	"os"
	"strings"
	"time"
)

var db *leveldb.DB
var wo *opt.Options
var ro *opt.ReadOptions
var json = jsoniter.ConfigCompatibleWithStandardLibrary
var producer *kafka.Producer
var consumer *kafka.Consumer
var (
	app        *cli.App
	configPath = "config.yml"
	topic      = "message_1"
	dbpath     = flag.String("dbpath", "/pdata/leveldb/user1", " leveldb local path")
	//dbpath = flag.String("dbpath", "/dianyi/data/leveldb/user.dat", " leveldb local path")
)

func workInit() {
	cnf, _ := loadConfig()
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cnf.Broker,
		"group.id":          "pushworker",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		fmt.Println("start error")
		panic(err)
	}
	consumer = c
}
func sendInit() {
	cnf, _ := loadConfig()
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": cnf.Broker, "queue.buffering.max.kbytes": 2000000,
		"queue.buffering.max.messages": 9000000})
	if err == nil {
		producer = p
	}
	wo = &opt.Options{
		ReadOnly: true,
	}
	tmpdb, err := leveldb.OpenFile(*dbpath, wo)
	if err != nil {

	} else {
		logging.Println(tmpdb.GetSnapshot())
	}
	db = tmpdb
}
func init() {

	//defer db.Close()
	//Working = false
	//common.SysCron.AddFunc("0 0/1 * * * *", Work)
	// Initialise a CLI app
	app = cli.NewApp()
	app.Name = "machinery"
	app.Usage = "machinery worker and send example tasks with machinery send"
	app.Author = "Richard Knop"
	app.Email = "risoknop@gmail.com"
	app.Version = "0.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "c",
			Value:       "",
			Destination: &configPath,
			Usage:       "Path to a configuration file",
		},
	}
}

func main() {
	// Set the CLI app commands
	app.Commands = []cli.Command{
		{
			Name:  "worker",
			Usage: "launch machinery worker",
			Action: func(c *cli.Context) error {
				workInit()
				common.InitWorkDomain()
				if err := worker(); err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				return nil
			},
		},
		{
			Name:  "send",
			Usage: "send example tasks ",
			Action: func(c *cli.Context) error {
				//if err := Work(); err != nil {
				//	return cli.NewExitError(err.Error(), 1)
				//}
				common.InitSend()
				sendInit()
				deliveryChan := make(chan kafka.Event)
				go func() {
					for {
						select {
						case rec, ok := <-deliveryChan:
							if !ok {
								logging.Errorln("consumer msg fail" + rec.String())
							}

						}
					}
				}()
				Work(&deliveryChan)
				return nil
			},
		},
	}
	os.Args = append(os.Args, "worker")
	// Run the CLI app
	app.Run(os.Args)
}

func loadConfig() (*config.Config, error) {
	return config.NewFromYaml("config.yml", true)
}

func worker() error {
	fmt.Println("start consumer")
	consumer.SubscribeTopics([]string{topic}, nil)
	p := pool.NewLimited(2000)
	batch := p.Batch()
	i := 0
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			//fmt.Printf(string(msg.Value))
			batch.Queue(task.Push(strings.Split(string(msg.Value), "+_+")))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
		i++
		if i == 5000 {
			ch := make(chan int, 1)
			go func() {
				logging.Infoln("send 5000 start")
				batch.QueueComplete()
				batch.WaitAll()
				ch <- 1
			}()

			select {
			case res := <-ch:
				logging.Infoln(fmt.Sprintf("send 5000 finish,result %d", res))
			case <-time.After(time.Second * 10):
				logging.Infoln("send 5000 timeout")
				batch.Cancel()
			}

			//batch.Cancel()
			//for email := range batch.Results() {
			//	if err := email.Error(); err != nil {
			//		logging.Errorln(err)
			//	}
			//	logging.Infoln(fmt.Sprintf("user %s send success",email.Value()))
			//}
			batch = p.Batch()

			i = 0
		}
	}
	consumer.Close()
	return nil
}
func Work(deliveryChan *chan kafka.Event) {

	it := db.NewIterator(nil, ro)
	//defer it.Close()
	defer it.Release()

	last := metrics.GetLastKey()
	it.First()
	if last == nil || *last == string(it.Key()) {
		logging.Infoln("now is first,skip to last")
		it.Last()
	} else {
		it.Seek([]byte(*last))
	}

	//// Decode subscription
	//s := &webpush.Subscription{
	//	Endpoint: user[9],
	//	Keys: webpush.Keys{
	//		P256dh: user[16],
	//		Auth:   user[34],
	//	},
	//}
	//if domain, ok := common.Domains[user[44]]; ok {
	//	// Send Notification
	//	start := time.Now()
	//	res, err := webpush.SendNotification([]byte(msg), s, &webpush.Options{
	//		Subscriber:      user[34],
	//		VAPIDPublicKey:  *domain.PublicKey,
	//		VAPIDPrivateKey: *domain.PrivateKey,
	//		TTL:             30,
	//	})
	//	if err != nil {
	//		logging.Errorln(err)
	//	}
	//	//campaign := common.Campaigns[*campaignId]
	//	//if res.StatusCode == 201 {
	//	//	go metrics.SendSuccess(campaign.OfferId, campaignId)
	//	//} else {
	//	//	go metrics.SendFail(campaign.OfferId, campaignId)
	//	//}
	//
	//	return fmt.Sprintf("send to user %s result:%d,spend %d ms", user[0], res.StatusCode, time.Since(start)/time.Millisecond), nil // everything ok, send nil, error if not
	//} else {
	//	return fmt.Sprintf("send to user %s fail:miss domain %s", user[0], user[44]), nil // everything ok, send nil, error if not
	//}

	for ; it.Valid(); it.Prev() {
		record := strings.Split(string(it.Value()), "+_+")
		for _, v := range common.Campaigns {
			//if arrays.ContainsString(v.Target.CountryExclude, record[4]) == -1 || arrays.ContainsString(v.Target.CountryInclude, record[4]) > -1 {
			//	//batch.Queue(sendMsg(record, base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s~%s~%s", "", record[0], *v.CampaignId))), v.CampaignId))
			//	//user:= [6]string{record[9],record[16],record[34],record[44],base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s~%s~%s", "", record[0], *v.CampaignId))),record[0]}
			//	//user1 := fmt.Sprintf("%s+_+%s+_+%s+_+%s+_+%s+_+%s", user[9], user[16], user[34], user[44], "MSG", user[0])
			//	//user:=fmt.Sprintf("%s+_+%s+_+%s+_+%s+_+%s+_+%s",record[9],record[16],record[34],record[44],base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s~%s~%s", "", record[0], *v.CampaignId))),record[0])
			//	user:=strings.ReplaceAll(string(it.Value()),"+_+MSG+_+",base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s~%s~%s", "", record[5], *v.CampaignId))))
			//	send(&user,deliveryChan)
			//}
			user := strings.ReplaceAll(string(it.Value()), "+_+MSG+_+", fmt.Sprintf("+_+%s+_+", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s~%s~%s", "", record[5], *v.CampaignId)))))
			//logging.Infoln(fmt.Sprintf("send msg:%s",user))
			send(&user, deliveryChan)
		}

	}
	Work(deliveryChan)
}

func send(user *string, deliveryChan *chan kafka.Event) {
	err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(*user),
		Headers:        []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary")}},
	}, *deliveryChan)
	if err != nil {
		fmt.Println(err.Error())
	}
}
