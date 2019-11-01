package task

import (
	"encoding/base64"
	"fmt"
	"github.com/SherClockHolmes/webpush-go"
	"github.com/cgiser/pushworker/common"
	"github.com/cgiser/pushworker/logging"
)

// Add ...
func Push(args ...[]string) (string, error) {

	return sendMsg(args[0])
}

func sendMsg(user []string) (string, error) {
	// Decode subscription
	s := &webpush.Subscription{
		Endpoint: user[0],
		Keys: webpush.Keys{
			P256dh: user[1],
			Auth:   user[2],
		},
	}
	x, _ := base64.StdEncoding.DecodeString(user[4])
	fmt.Println(string(x))
	if domain, ok := common.Domains[user[3]]; ok {
		// Send Notification
		res, err := webpush.SendNotification([]byte(user[4]), s, &webpush.Options{
			Subscriber:      user[2],
			VAPIDPublicKey:  *domain.PublicKey,
			VAPIDPrivateKey: *domain.PrivateKey,
			TTL:             30,
		})
		if err != nil {
			logging.Errorln(err)
		}
		//campaign := common.Campaigns[*campaignId]
		//if res.StatusCode == 201 {
		//	go metrics.SendSuccess(campaign.OfferId, campaignId)
		//} else {
		//	go metrics.SendFail(campaign.OfferId, campaignId)
		//}

		return fmt.Sprintf("send to user %s result:%s", user[5], res.Status), nil // everything ok, send nil, error if not
	} else {
		return fmt.Sprintf("send to user %s fail:miss domain %s", user[5], user[3]), nil // everything ok, send nil, error if not
	}
	//msg := base64.StdEncoding.EncodeToString([]byte("1~2~30001"))
	//s := &webpush.Subscription{
	//	Endpoint: "https://fcm.googleapis.com/fcm/send/eKijXBDqQlk:APA91bGBsdSVvQ9PN5FcDKX1aY8ATQSV7gM3TBCH8_AEyVeI7defMtgxZ227b-c80J8CY0GNe-psPRGOQNA1SdBZ9b4uDoft3Y4x4gxnenQ1-47d7IhbT5NUma9WXd5uTi2ZoP-vo-ur",
	//	Keys: webpush.Keys{
	//		P256dh: "BBKAgZs1FU7pU+i/WOY7RCKBnZ159tKKvrjROdYrqE7Ysw00PnXp3/nRnChLLE4kghTrz0vhjjpAF8zmNGXfYuc=",
	//		Auth:   "RzcdAcNb1RiJ3pKJ1ujwzw==",
	//	},
	//}
	//
	//// Send Notification
	//res, _ := webpush.SendNotification([]byte(msg), s, &webpush.Options{
	//	Subscriber:      "RzcdAcNb1RiJ3pKJ1ujwzw==",
	//	VAPIDPublicKey:  "BJgwoxs8qBkc2TG3WHkAr7ESXjmCZgEnro3lk0V9fcYiA6dmth25BZZusMwQKGK0dF8cUqy_F_LR4cOT2Bu9Qgg",
	//	VAPIDPrivateKey: "jO9RsCsY6uejxetA1OyzWc5SWmv0GbQVoBt1lvB0QbU",
	//	TTL:             30,
	//})
	//return fmt.Sprintf("send to user %s result:%s", user[5], res.Status), nil // everything ok, send nil, error if not
}
