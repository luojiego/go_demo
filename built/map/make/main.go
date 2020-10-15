package main

import "fmt"

type Data struct {
	Count int
	M     map[int]int
}

func initMap() *Data {
	return &Data{
		M: make(map[int]int),
	}
}

func test(d *Data) {
	fmt.Println("test: ", &d)
	d = initMap()
	fmt.Println("test after init: ", &d)
	d.Count = 987
	d.M[1] = 1000
	d.M[2] = 2000
}

func main() {
	d1 := Data{}
	fmt.Printf("main d1: %p\n", &d1)
	d := &d1
	fmt.Println("main d: ", &d)
	test(d)
	fmt.Println(d)
}
