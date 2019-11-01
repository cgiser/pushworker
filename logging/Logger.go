package logging

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var (
	Logger *logrus.Logger
)

func Println(args ...interface{}) {
	Logger.Infoln(args)
}
func Printf(format string, args ...interface{}) {
	Logger.Infoln(format, args)
}
func Infoln(args ...interface{}) {
	Logger.Infoln(args)
}
func Infof(format string, args ...interface{}) {
	Logger.Infoln(format, args)
}

func Errorln(args ...interface{}) {
	Logger.Errorln(args)
}
func Errorf(format string, args ...interface{}) {
	Logger.Errorf(format, args)
}
func init() {
	Logger = logrus.New()
	Logger.Formatter = &logrus.TextFormatter{}
	Logger.SetLevel(logrus.InfoLevel)
	ll := &lumberjack.Logger{
		Filename:   "/dianyi/log/api.log",
		MaxSize:    128, // megabytes
		MaxBackups: 5,
		MaxAge:     28,   // days
		Compress:   true, // disabled by default
	}
	Logger.Out = ll
}
