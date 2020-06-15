package main

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	_ "github.com/micro/go-plugins/broker/kafka/v2"
	"log"
	"time"
)

var (
	topic  string = "go-micro-kafka"
	send_i        = 0
	recv_i        = 0
)

func pub(brk broker.Broker) {
	for range time.Tick(time.Second) {
		msg := &broker.Message{
			Header: map[string]string{"id" : fmt.Sprintf("%d", send_i)},
			Body:   []byte(fmt.Sprintf("%d:%s", send_i, time.Now().String())),
		}
		if err := brk.Publish(topic,msg); err != nil {
			log.Printf("[pub] failed: %v\n", err)
		} else  {
			log.Printf("[pub] pubbed message:%s\n", string(msg.Body))
		}
		send_i++
	}
}

func sub(brk broker.Broker)  {
	_, err := brk.Subscribe(topic, func(event broker.Event) error {
		recv_i++
		fmt.Println("[sub] received message:", string(event.Message().Body), ", header:", event.Message().Header)
		return nil
	}, broker.Queue(topic))
	if err != nil {
		fmt.Println(err)
	}
}

func main()  {
	service := micro.NewService(
		micro.Name("com.foo.broker.example"),
	)
	service.Init(micro.AfterStart(func() error {
		brk := service.Options().Broker
		if err := brk.Connect(); err != nil {
			log.Fatalf("Broker connect error:%v", err)
		}
		go sub(brk)
		go pub(brk)
		return nil
	}),
		micro.BeforeStop(func() error {
			log.Printf("send:%d,recv:%d\n", send_i, recv_i)
			return nil
		}))

	service.Run()
}
