package main

import (
	"github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"path"
	"time"
)

func test() {
	log.WithFields(log.Fields{
		"func": test,
	}).Info("test func")
}

func init() {
	baseLogPath := path.Join("./log", "rotate.log")
	writer, _ := rotatelogs.New(
		baseLogPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPath),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		//rotatelogs.WithRotationCount(10),
		rotatelogs.WithRotationTime(1*time.Minute),
	)
	//log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(writer)
	log.SetLevel(log.TraceLevel)
}

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	log.Info("hello")

	l := log.WithFields(log.Fields{"name": "Roger"})

	for {
		l.Debug("你说啥，我听不见")
		l.Debug("我再说一遍，你好吗？")
		l.Println("hello")
		time.Sleep(time.Second)
	}

}
