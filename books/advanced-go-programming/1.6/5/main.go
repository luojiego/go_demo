package main

import (
	"fmt"
	"runtime"
)

func SearchByBing(src string) string {
	return "Bing: " + src
}

func SearchByGoogle(src string) string {
	return "Google: " + src
}

func SearchByBaidu(src string) string {
	return "Baidu: " + src
}

func main() {
	max := runtime.GOMAXPROCS(1)
	fmt.Println(max)
	ch := make(chan string, 32)

	go func() {
		ch <- SearchByBing("golang")
	}()

	go func() {
		ch <- SearchByGoogle("golang")
	}()

	go func() {
		ch <- SearchByBaidu("golang")
	}()

	fmt.Println(<-ch)
}
