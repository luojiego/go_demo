package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func test() {
	logrus.WithFields(logrus.Fields{
		"func": test,
	}).Info("test func")
}

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{DisableTimestamp: true})
	//logrus.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true})
	file, err := os.Create("try.log")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	logrus.SetOutput(file)
	//logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.TraceLevel)
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	l := logrus.WithFields(logrus.Fields{"name": "Roger"})
	l.Debug("你走")
	l.Debug("你快走")
	l.Println("hello")

	test()
	//fmt.Println(l.Buffer.String())
	//l.Log()
}
