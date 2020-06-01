package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"io/ioutil"
	"os"
	"time"
)

var (
	buf []byte
)

func init() {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	buf, err = ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
}

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

	n := time.Now()
	for i := 0; i < 10000; i++ {
		//stringTest(c) 4.96s
		//hSet(c) 4.55432s
		//hMSet(c) 4.63374s
		//hSetLarge(c) 43.17409s
	}
	fmt.Printf("%0.5fs\n", time.Since(n).Seconds())
}

func stringTest(c *redis.Client) {
	_, err := c.Set("name", "Roger", -1).Result()
	if err != nil {
		panic(err)
	}

	_, err = c.Get("name").Result()
	if err != nil {
		panic(err)
	}

}

func hSet(c *redis.Client) {
	_, err := c.HSet("data", "name", "Roger").Result()
	if err != nil {
		panic(err)
	}

	_, err = c.HGet("data", "name").Result()
	if err != nil {
		panic(err)
	}
}

func hSetLarge(c *redis.Client) {
	_, err := c.HSet("data", "info", buf).Result()
	if err != nil {
		panic(err)
	}

	_, err = c.HGet("data", "info").Result()
	if err != nil {
		panic(err)
	}
}

func HMSet(c *redis.Client) {

	f := map[string]interface{}{
		"name": "Roger",
		"age":  30,
	}
	_, err := c.HMSet("data", f).Result()
	if err != nil {
		panic(err)
	}

	_, err = c.HMGet("data", []string{"name", "age"}...).Result()
	if err != nil {
		panic(err)
	}
}
