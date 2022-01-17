package main

import (
	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
)

func main() {
	hook := lumberjack.Logger{
		Filename:   "./logs/logrus-lumberjack.log", // 日志文件路径
		MaxSize:    10,                             // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 3,                              // 日志文件最多保存多少个备份
		MaxAge:     14,                             // 文件最多保存多少天
		Compress:   false,                          // 是否压缩
	}

	log.SetOutput(&hook)
	log.SetLevel(log.TraceLevel)
	log.SetFormatter(&log.JSONFormatter{})
	//log.SetReportCaller(true)

	l := log.WithFields(log.Fields{
		"animal": "walrus",
		"name":   "二宝宝",
	})

	l.Info("A walrus appears")

	l1 := log.WithFields(log.Fields{
		"name": "Roger",
	})

	l1.Info("hello")

	l2 := l.WithFields(log.Fields{"id": 10000, "name": "刘延军"})
	l2.Info("haha ")
	//l := log.WithFields(log.Fields{"name": "Roger"})
	return
	index := 0
	for {
		l.Debug("你说啥，我听不见")
		l.Debug("我再说一遍，你好吗？")
		l.Println("hello")
		index++
		if index > 200000 {
			break
		}
	}
}
