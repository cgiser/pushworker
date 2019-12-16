package model

type ImpressionParam struct {
	MsgId      string `json:"msgId"`
	Uid        string `json:"uid"`
	CampaignId string `json:"cid"`
	//Pub  string  `json:"pub"`
	//MsgDetailsCoder  string  `json:"msgDetailsCoder"`
	//MsgDetailsId  string  `json:"msgDetailsId"`

	//String msgId = null;
	//String uid = null;
	//String pub = "0";
	//String msgDetailsCoder = null;
	//String msgDetailsId = null;
}

//{
//"//": "Visual Options",
//"body": "<String>",
//"icon": "<URL String>",
//"image": "<URL String>",
//"badge": "<URL String>",
//"vibrate": "<Array of Integers>",
//"sound": "<URL String>",
//"dir": "<String of 'auto' | 'ltr' | 'rtl'>",
//
//"//": "Behavioral Options",
//"tag": "<String>",
//"data": "<Anything>",
//"requireInteraction": "<boolean>",
//"renotify": "<Boolean>",
//"silent": "<Boolean>",
//
//"//": "Both visual & behavioral options",
//"actions": "<Array of Strings>",
//
//"//": "Information Option. No visual affect.",
//"timestamp": "<Long>"
//}
type Msg struct {
	Title              *string                `json:"title",bson:"title"`
	Body               *string                `json:"body"`
	Icon               *string                `json:"icon"`
	Image              *string                `json:"image"`
	Badge              *string                `json:"badge"`
	Vibrate            []int                  `json:"vibrate"`
	Sound              *string                `json:"sound"`
	Dir                *string                `json:"dir"`
	Tag                *string                `json:"tag"`
	Data               map[string]interface{} `json:"data"`
	RequireInteraction bool                   `json:"requireInteraction"`
	Renotify           bool                   `json:"renotify"`
	Silent             bool                   `json:"silent"`
	Actions            []*Action              `json:"actions"`
	Timestamp          int64                  `json:"timestamp"`
	Url                *string                `json:"url"`
	CampaignId         *string                `json:"cid"`
	Cpc                *float64               `json:"cpc"`
	Tracker            *string                `json:"tracker"`
}

type Action struct {
	Icon   *string `bson:"icon" json:"icon"`
	Title  *string `bson:"title" json:"title"`
	Action *string `bson:"action" json:"action"`
}

type Offer struct {
	OfferId *string  `bson:"offerId"`
	Name    *string  `bson:"name"`
	Url     *string  `bson:"url"`
	Price   *float64 `bson:"price"`
	Title   *string  `bson:"title"`
	Body    *string  `bson:"body"`
	Icon    *string  `bson:"icon"`
	Image   *string  `bson:"image"`
	Badge   *string  `bson:"badge"`
}

type Target struct {
	CountryInclude []string `bson:"country_include"`
	CountryExclude []string `bson:"country_exclude"`
	User           *string  `bson:"user"`
}

type Campaign struct {
	CampaignId *string   `bson:"campaignId"`
	Name       *string   `bson:"name"`
	Title      *string   `bson:"title"`
	Body       *string   `bson:"body"`
	Icon       *string   `bson:"icon"`
	Image      *string   `bson:"image"`
	Badge      *string   `bson:"badge"`
	OfferId    *string   `bson:"offerId"`
	Target     *Target   `bson:"target"`
	Action     []*Action `bson:"action"`
	Cpc        *float64  `json:"cpc"`
}

type GktUser struct {
	Token *string `bson:"token"`
}

//"title" : "📢 Binabati kita !!",
//"body" : "Ang PHP 100000.00 ay nasa iyong account, gamitin ito ngayon! 💰💰💰",
//"image" : null,
//"icon" : "https://images.pushuse.com/images/34888c2a-1564-43c1-a49b-2669633a08e9.png",
//"url" : "https://fedifice-prionald.com/a2021764-4fb9-4265-9eff-144ba4e542af?msgId={msgId}",
//"badge" : "https://images.pushuse.com/images/d2eb91cd-2025-43c5-9943-e756c835ff62.png",
//"countryCode" : "PH",
//"createTime" : "16/6/2019 15:10:14",
//"createUtcTime" : "1560697814147",
//"day" : "16/6/2019",
//"dayAndHour" : "2019-06-17_10",
//"defaultMessage" : "0",
//"deleteFlag" : "0",
//"excludeCountryCode" : null,
//"hour" : "15",
//"close" : "0",
//"autoClose" : "0",
//"msgId" : "36257360670a4188a09cbf2a1ecd74ba",
//"msgType" : "DEFAULT",
//"sendType" : null,
//"setFlag" : "0",
//"bodyTranslateFlag" : "0",
//"clearEvent" : null,
//"clearFlag" : "0",
//"subscribedDomain" : null,
//"tag" : "Fastloan",
//"tagClearFlag" : "0",
//"timeStage" : null,
//"titleTranslateFlag" : "0",
//"writeRedisFlag" : null,
//"updateTime" : "16/6/2019 15:10:14",
//"updateUtcTime" : "1560697814147",
//"userName" : "amy.yang",
//"vibrate" : "[3000]",
//"actions" : "[{\"icon\":\"\",\"action\":\"0\",\"title\":\"Mag-click ngayon\"},{\"icon\":\"\",\"action\":\"1\",\"title\":\"Huwag pansinin\"}]",
//"clearCondition" : "{}",
//"ratio" : null,
//"cron" : null,
//"sendTime" : null,
//"beginTime" : null,
//"lastMinutes" : null,
//"sendIntervalMinutes" : null
type Domain struct {
	DomainId   *string `bson:"domainId"`
	PrivateKey *string `bson:"privateKey"`
	PublicKey  *string `bson:"publicKey"`
	Url        *string `bson:"url"`
}

//"createDate" : null,
//"domainId" : "9e48f174a20c4c40a479f20eaaa070f6",
//"privateKey" : "7jCleNiI9S0rtkjgwZuz_yWHBP7UccDglPkzmPOjarY",
//"publicKey" : "BIFDSTB_MqtWxMXpPuyG3l6A5-B-UNTF4wCrq1kRZIV3rp72MvOKNgDdQz4X-DPfKa7Ls4NvmtfpiVM4uoNGcJk",
//"url" : "news-subscribe.com"
type User struct {
	Uid                 *string `bson:"uid"`
	CountryCode         *string `bson:"countryCode"`
	Endpoint            *string `bson:"endpoint"`
	Key                 *string `bson:"key"`
	Token               *string `bson:"token"`
	SubscribedTopDomain *string `bson:"subscribedTopDomain"`
}

//{"endpoint":"https://fcm.googleapis.com/fcm/send/ezIYfyHRUbs:APA91bEEZiEisLUWa2SS1CUioPB7n7G8_WJDloTw8FqZa-tjjDhgqu_TMrFZ-YHVWl4e1wXtnMRyhOVmjzndgIzvf3t1n9QbxXd71pFcSkHGJCKP9Hb8RAyIfBjtI4khAkg7AFfOntH1","expirationTime":null,"keys":{"p256dh":"BHzMOdt-saScBs1Bp4EhnKZfNE8j0WMCKrHGzsOAIGxTKFfkWUAHz_UK3pIC28fUL75XbcLu2D9In7PgtEa64VI","auth":"5qcFUezUIbLychtx5T3j3g"}}
type SubscribeUser struct {
	Uid              *string            `bson:"uid"`
	Endpoint         *string            `bson:"endpoint"`
	Keys             *SubscribeUserKeys `bson:"keys"`
	Refer            *string            `bson:"refer"`
	Ua               *string            `bson:"ua"`
	Ip               *string            `bson:"ip"`
	SubscribedDomain *string            `bson:"subscribedDomain"`
}

type SubscribeUserKeys struct {
	P256dh *string `bson:"p256dh"`
	Auth   *string `bson:"auth"`
}
