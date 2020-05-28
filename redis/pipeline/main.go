package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
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
	fmt.Println("ping: ", result)
	/*m := map[string]*redis.StringCmd{}
	pipeline := c.Pipeline()
	m["a"] = pipeline.Get("a")
	m["b"] = pipeline.Get("b")*/

	pipeline := c.Pipeline()
	pipeline.HSet("10001", "name", "Roger")
	pipeline.Expire("10001", 150*time.Second)
	cmders, err := pipeline.Exec()
	if err != nil {
		panic("exec: " + err.Error())
	}

	for _, exec := range cmders {
		cmd := exec.(*redis.BoolCmd)
		s, err := cmd.Result()
		if err != nil {
			panic(err)
		}
		fmt.Println(s)
	}

}
