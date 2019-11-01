package metrics

import (
	"flag"
	"github.com/cgiser/pushworker/logging"
	jsoniter "github.com/json-iterator/go"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

var db *leveldb.DB
var oo *opt.Options
var wo *opt.WriteOptions
var ro *opt.ReadOptions
var json = jsoniter.ConfigCompatibleWithStandardLibrary
var (
	keypath = flag.String("keypath", "/pdata/leveldb/workerkey", " leveldb local path")
)

func init() {
	wo = &opt.WriteOptions{}
	db, _ = leveldb.OpenFile(*keypath, oo)
}
func GetLastKey() *string {
	//session := common.GetMongoSession()
	//defer session.Close()
	//campaignstats := session.DB(common.MogDB).C("lastkey")
	//query := bson.M{
	//	"key":"lastkey",
	//}
	//
	//last:=model.LastKey{}
	//err:=campaignstats.Find(query).One(&last)
	//if err!=nil{
	//	logging.Errorln(err)
	//	return nil
	//}else {
	//	return last.Value
	//}
	val, err := db.Get([]byte("lastkey"), ro)
	key := string(val)
	if err != nil {
		logging.Errorln(err)
		return nil
	} else {
		return &key
	}
}
