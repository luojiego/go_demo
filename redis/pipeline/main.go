package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	c := redis.NewClient(
		&redis.Options{
			Addr: "192.168.196.50:6011",
			DB:   15,
		})
	result, err := c.Ping().Result()
	if err != nil {
		panic("ping: " + err.Error())
	}
	fmt.Println(result)
	pipeline := c.Pipeline()
	pipeline.Get("A")
	pipeline.Get("B")

	cmders, err := pipeline.Exec()
	if err != nil {
		panic("exec: " + err.Error())
	}

	for _, exec := range cmders {
		cmd := exec.(*redis.StringCmd)
		s, err := cmd.Result()
		if err != nil {
			panic(err)
		}
		fmt.Println(s)
	}

}
