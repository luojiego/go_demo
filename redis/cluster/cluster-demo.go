package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"192.168.196.50:6020",
			"192.168.196.50:6021",
			"192.168.196.50:6022",
		},
	})
	defer client.Close()

	result, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

	userId := "1000004"
	s, err := client.HGet(userId, "game_data").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("can't find ", userId)
		} else {
			panic(err)
		}
	}

	fmt.Println(s, err)

	b, err := client.HSet("10007", "name", "luojie").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(b, err)

}
