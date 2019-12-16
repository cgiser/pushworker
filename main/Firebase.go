package main

import (
	"encoding/base64"
	"fmt"
	"github.com/appleboy/go-fcm"
	"github.com/cgiser/pushworker/common"
	"github.com/cgiser/pushworker/logging"
	"github.com/cgiser/pushworker/model"
	"log"
)

var client *fcm.Client

func init() {
	// init client
	c, err := fcm.NewClient("AIzaSyA99sBKvrpcUTAKh4YJWQDf5O9NHn3rVs8")
	if err != nil {
		log.Fatalln(err)
	}
	client = c
	common.InitSend()
}
func work() {
	logging.Println("开始发送campaign")
	session := common.GetMongoSession()
	defer session.Close()
	bundle := session.DB(common.MogDB).C("gkt_user")

	iter := bundle.Find(nil).Iter()

	user := model.GktUser{}
	for iter.Next(&user) {
		for _, v := range common.Campaigns {
			if v.Target != nil && v.Target.User != nil && *v.Target.User == "gkt" {
				push(&user, v)
			}
		}

	}
	work()
}
func push(user *model.GktUser, cammpaign *model.Campaign) {

	// You can use your HTTPClient
	//client.SetHTTPClient(client)

	//data := map[string]interface{}{
	//	"messageId": "MX4yfjQwMDAx",
	//	"title": "Sexy Video",
	//	"icon":"https://scontent-ort2-1.cdninstagram.com/vp/4dec53280540b489ec6b370ea168c84b/5E3425D3/t51.2885-19/s150x150/54512495_2182097381837512_5501368273087758336_n.jpg?_nc_ht=scontent-ort2-1.cdninstagram.com",
	//	"body":"Woman's secret video！",
	//}

	// Create the message to be sent.
	//msg := &fcm.Message{
	//	To: "e8shFm2Y5NI:APA91bG5suN8l-Edmo10tQLr32oSiDM11CU5P-3NM4SwdPTYNI9zm8AB2AHabID3mfN_wM0L1sIKy3Vuxwb51jJ8hINtOLx52Z7hY4H4fCj1yxYz-mjoi_uhOPOOdkXzo9LRzdKJFf1A",
	//	Data: map[string]interface{}{
	//			"messageId": "MX4yfjQwMDAx",
	//			//"title": "Sexy Video",
	//			//"icon":"https://scontent-ort2-1.cdninstagram.com/vp/4dec53280540b489ec6b370ea168c84b/5E3425D3/t51.2885-19/s150x150/54512495_2182097381837512_5501368273087758336_n.jpg?_nc_ht=scontent-ort2-1.cdninstagram.com",
	//			//"body":"Woman's secret video！",
	//	},
	//	Notification:&fcm.Notification{
	//		Title:"Sexy Video",
	//		Body:"Woman's secret video！",
	//		Icon:"https://scontent-ort2-1.cdninstagram.com/vp/4dec53280540b489ec6b370ea168c84b/5E3425D3/t51.2885-19/s150x150/54512495_2182097381837512_5501368273087758336_n.jpg?_nc_ht=scontent-ort2-1.cdninstagram.com",
	//
	//	},
	//}
	messageId := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("1~2~%s~gkt", *cammpaign.CampaignId)))
	msg := &fcm.Message{
		To: *user.Token,
		Data: map[string]interface{}{
			"messageId": messageId,
			//"title": "Sexy Video",
			//"icon":"https://scontent-ort2-1.cdninstagram.com/vp/4dec53280540b489ec6b370ea168c84b/5E3425D3/t51.2885-19/s150x150/54512495_2182097381837512_5501368273087758336_n.jpg?_nc_ht=scontent-ort2-1.cdninstagram.com",
			//"body":"Woman's secret video！",
		},
		Notification: &fcm.Notification{
			Title: *cammpaign.Title,
			Body:  *cammpaign.Body,
			Icon:  *cammpaign.Image,
		},
	}

	// Create a FCM client to send the message.

	// Send the message and receive the response without retries.
	response, err := client.Send(msg)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", response)
}
func main() {
	work()
}
