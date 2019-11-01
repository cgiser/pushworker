package common

import (
	"encoding/json"
	"fmt"
	"github.com/cgiser/pushworker/logging"
	"github.com/cgiser/pushworker/model"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/robfig/cron"
)

var NowVersion = "v0.1"

var (
	MogUrl    string
	MogDB     string
	Offers    = map[string]*model.Offer{}
	Campaigns = map[string]*model.Campaign{}
	Domains   = map[string]*model.Domain{}
	SysCron   = cron.New()
)
var (
	mgoBidSession *mgo.Session
	domainList    = "[{\"privateKey\":\"7jCleNiI9S0rtkjgwZuz_yWHBP7UccDglPkzmPOjarY\",\"publicKey\":\"BIFDSTB_MqtWxMXpPuyG3l6A5-B-UNTF4wCrq1kRZIV3rp72MvOKNgDdQz4X-DPfKa7Ls4NvmtfpiVM4uoNGcJk\",\"url\":\"news-subscribe.com\"},{\"privateKey\":\"tv8XG5utUx0SEWG9YnAlRy7SlIklfghUz9f6YAo3nXk\",\"publicKey\":\"BHkHqWXVvumsjNeiIgbWO5EAanG5O3INgcj_Nr4jOA_VV1kX_DUmTtEmgt38H_-xCbi4EyDMifqDwHnXNESblfs\",\"url\":\"savingmoneys.com\"},{\"privateKey\":\"cRmaNU-dNKV10m5JguCQHvUMzzwunJnYs9-eGvXiQTc\",\"publicKey\":\"BCHNSTyGwA-d38mn5iRdyowxRz8tPQm9zNJOd77iPRcd_uT_cf4yUe_JKTT0OE996VB8yaXkIiBF47spcskzj1s\",\"url\":\"supermarketsaves.com\"},{\"privateKey\":\"FYo8QLdkIqwn66i3Vu93Z-D63uEO4GWPKBslUCZrGW8\",\"publicKey\":\"BO-n1Xqz_EKCZGCirK93gg8QQjea2mQP8NKA_56AAZCGqnwjTZtLTFonBi4MPjQOflg3M7ZwjmblGy0_hYHnszY\",\"url\":\"supermarket-coupons.info\"},{\"privateKey\":\"LN2xQkIVRK19lgTavlLDqqZzsEmYzvKsAjfiz4Vw-dM\",\"publicKey\":\"BIGTYckhgvBC4ZasviS3Z32tGzKhH-d7CYLCdJg5Sj-XdKCKH_1li7vjqzwkexg5ozAqWUK9j0mLOuqclBJgO24\",\"url\":\"walmartgiftcard.vip\"},{\"privateKey\":\"-AcyWvurWJOh4aD_l8y3960j_3J7iwxdYsxe90DLsLo\",\"publicKey\":\"BNsQz5ZbyajRTM7PKBGT7rbU4IUnZovJKBeFjNDuCnVgNjI_oNNDgkz6h5s8Hf0ueMvspVqxMqvgJY2-DBmY0n8\",\"url\":\"savingmoneycard.com\"},{\"privateKey\":\"NRn0J6FSWuQQrD2d8zeIJqC5AorlMtJ-K_C6YA3ZYCY\",\"publicKey\":\"BOvxoFbZcphU_Y-GjJWIyziVu43nqm72uyr4EpfMsAzElg5OzcMI-GLg7QP5jINzPXNFI5hdlRsqXUu5Lp34qYI\",\"url\":\"check-now.net\"},{\"privateKey\":\"Xcum2KlWZ8RBGtI6P2ikn2eGT6PQ57ZvT7zcpV_gy3A\",\"publicKey\":\"BMdDlcSC-p-CZ8-pqmrJ3uix-QTantR5NTo9nhQJa92pmman_wqEsnXu5rD206bYH8H9Bqdv10pw8xyjYxnGGS0\",\"url\":\"yahoo-news.co\"},{\"privateKey\":\"iFoZjQZS-dV8EpGdKkdrzenrpaMDZ1jngd5RQ_pQIJk\",\"publicKey\":\"BP0O0_1R4ZQclUsk2B0b0PMwah68oqd5xSrRmAaq3s1jnrn0b0_kKTZMYoKeviN4rS3P3p-WwWM1JcIz62h8q1I\",\"url\":\"googlo.co\"},{\"privateKey\":\"Uzko8SVjezLzd5UKQ7Yn39Ug7uaTYgCrne9tiFW7u9s\",\"publicKey\":\"BPwASBd6fZSe4lTeoYkCA8sJ66VFYWG6q11OL8C2p1CQZ0fbwScloPn5Y5pX6si1AG4WjDUVja-KZNvAQMIZJSM\",\"url\":\"andriod.co\"},{\"privateKey\":\"yEPyb3niiMxCKi1MSmmJ-QTxmlMZ4luKV9iRPx2iUPg\",\"publicKey\":\"BDMlGF1Hs2WXjxPRXUF-YwNapIJ177A_61gkE9offM4507VuG8mgVde0riHTiVJQN-oY4X_obA9R14yjhL7jv6Q\",\"url\":\"facebook-info.co\"},{\"privateKey\":\"jO9RsCsY6uejxetA1OyzWc5SWmv0GbQVoBt1lvB0QbU\",\"publicKey\":\"BJgwoxs8qBkc2TG3WHkAr7ESXjmCZgEnro3lk0V9fcYiA6dmth25BZZusMwQKGK0dF8cUqy_F_LR4cOT2Bu9Qgg\",\"url\":\"yourtuber.info\"},{\"privateKey\":\"FmHTzGFT225E00x9nahKCDwwEq_48E-dArLQIL1zDD0\",\"publicKey\":\"BIahmNegQAK4XcmDBIV97Z3Ted3YT9fcOwfsCnvKibtQIGXYmm2ldUYavk47mZi5_zIADTJf-3RSky_avk6oomg\",\"url\":\"tops.video\"},{\"privateKey\":\"Rw_CTEQdoScSj9IPuJftDdJTVcHl9u6wdYl3UjuH_7s\",\"publicKey\":\"BNdr_HCqBOtb0FtXpsIG9I42UDDo_AZsd4J1nMJFgqiOyTJra_CNCFKxzQj0m6J66RM8IOLz4ryMOObd_GJh7EE\",\"url\":\"from-mobile.com\"},{\"privateKey\":\"ovsD-ghk7BOH26SIDxSgBTNanBaJoa_kaLmGox7Yiyo\",\"publicKey\":\"BCVg-TphAcuV_pCRCpzzUf786tXbykNmHmI1inym04_MR2hwemY4rpej7k249atxS8WVFKXb3bBSn5lcpzS2OhI\",\"url\":\"systempush.com\"},{\"privateKey\":\"llSPn3zQ17NTBrncsmoyWwblnDIEPkZy7sj8BZdtF04\",\"publicKey\":\"BPmXUe6POWr-WvYp3dJ5eOtXqxGLSSSb9_7xjqI_dAkP6p_VMNeTWrjDBtBmNPiPYd3IX0Rpa4bkdcByQd2YDM8\",\"url\":\"insystem.video\"}]"
)

func InitSend() {
	//MogUrl = "mongodb://pro_push:Q4M5tHSvUBWEN2t5WKNF@10.12.135.93,10.12.35.91,10.12.35.194:27017/push?maxPoolSize=100&authSource=admin"
	MogUrl = "mongodb://pro_push:Q4M5tHSvUBWEN2t5WKNF@52.91.43.99,34.227.206.79,54.210.24.215:27017/push?maxPoolSize=100&authSource=admin"
	MogDB = "push"
	InitDomain()
	InitOffer()
	InitCampaign()
	SysCron.AddFunc("0 0/1 * * * *", InitOffer)
	SysCron.AddFunc("0 0/1 * * * *", InitCampaign)
	SysCron.Start()
}
func InitWorkDomain() {
	logging.Println("开始加载domain")
	keys := make([]model.Domain, 0)
	err := json.Unmarshal([]byte(domainList), &keys)
	if err == nil {
		apps := map[string]*model.Domain{}
		for _, v := range keys {
			s := v
			apps[fmt.Sprintf("%s", *v.Url)] = &s
		}
		Domains = apps
	}

}
func InitOffer() {
	logging.Println("开始加载offer")
	session := GetMongoSession()
	defer session.Close()
	bundle := session.DB(MogDB).C("offer")
	query := bson.M{}

	iter := bundle.Find(query).Iter()

	app := model.Offer{}
	apps := map[string]*model.Offer{}
	for iter.Next(&app) {
		a := app
		apps[fmt.Sprintf("%s", *app.OfferId)] = &a

	}
	Offers = apps
}
func InitDomain() {
	logging.Println("开始加载domain")
	session := GetMongoSession()
	defer session.Close()
	bundle := session.DB(MogDB).C("domain")
	query := bson.M{}

	iter := bundle.Find(query).Iter()

	app := model.Domain{}
	apps := map[string]*model.Domain{}
	for iter.Next(&app) {
		a := app
		apps[fmt.Sprintf("%s", *app.Url)] = &a

	}
	Domains = apps
}

func InitCampaign() {
	logging.Println("开始加载campaign")
	session := GetMongoSession()
	defer session.Close()
	bundle := session.DB(MogDB).C("campaign")
	query := bson.M{
		"status": "Active",
	}

	iter := bundle.Find(query).Iter()

	app := model.Campaign{}
	apps := map[string]*model.Campaign{}
	for iter.Next(&app) {
		a := app
		apps[fmt.Sprintf("%s", *app.CampaignId)] = &a

	}
	Campaigns = apps
}

/**
 * 公共方法，获取session，如果存在则拷贝一份
 */
func GetMongoSession() *mgo.Session {
	if mgoBidSession == nil {
		var err error
		val := MogUrl
		mgoBidSession, err = mgo.Dial(val)
		mgo.SetStats(true)
		mgoBidSession.SetMode(mgo.Primary, true)
		if err != nil {
			panic(err) //直接终止程序运行
		}
	}
	session := mgoBidSession.Clone()
	//// Optional. Switch the session to a monotonic behavior.
	//session.SetMode(mgo.Eventual, true)
	return session
}
