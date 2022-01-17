package main

import "fmt"

func decorate(f func(s string)) func(s string) {
	return func(s string) {
		fmt.Println("started")
		f(s)
		fmt.Println("end")
	}
}

func Hello(s string) {
	fmt.Println(s)
}

func main() {
	decorate(Hello)("Hello world!")
}
