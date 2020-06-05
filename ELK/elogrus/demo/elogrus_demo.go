package main

import (
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"
	"os"
	"strconv"
	"time"
)

var (
	log *logrus.Logger
)

func main() {
	log = logrus.New()
	//log.SetFormatter(&logrus.JSONFormatter{})
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://192.168.196.50:9200"))
	if err != nil {
		log.Panicf("NewClient: ", err)
	}
	index := os.Getenv("INDEX")
	log.Info("index: ", index)
	hook, err := elogrus.NewElasticHook(client, "hostname", logrus.TraceLevel, index)
	if err != nil {
		log.Panic(err)
	}
	log.Hooks.Add(hook)

	/*log.WithFields(logrus.Fields{
		"name": "luojie",
		"age":  18,
	}).Info("ping")*/

	log.Info("the server shutdown...")
	test2()
	test3()
	//test()
}

func test2() {
	n := time.Now().Unix()
	log.WithFields(logrus.Fields{
		"name": "Roger" + strconv.FormatInt(n, 10),
	}).Info("你好")
	//time.Sleep(time.Second)
	log.WithFields(logrus.Fields{
		"name": "Roger" + strconv.FormatInt(n+1, 10),
	}).Info("hello world")
}

func test3() {
	n := time.Now().Unix()
	log.WithFields(logrus.Fields{
		"name": "test3" + strconv.FormatInt(n, 10),
	}).Info("你好")
	//time.Sleep(time.Second)
	log.WithFields(logrus.Fields{
		"name": "test3" + strconv.FormatInt(n+1, 10),
	}).Info("hello world")
	//time.Sleep(time.Second)
}

func test() {
	for i := 0; i < 100; i++ {
		log.WithFields(logrus.Fields{
			"name": "test",
		}).Infof("index: %d", i+1)
		log.WithFields(logrus.Fields{
			"name": "test",
		}).Infof("index: %d", i*100+1)
		time.Sleep(time.Second)
	}

}
