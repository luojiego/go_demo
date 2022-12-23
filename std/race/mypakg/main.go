//go:build !race
// +build !race

package main

import "fmt"

func main() {
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["1"] = "a" // First conflicting access.
		c <- true
	}()
	m["2"] = "b" // Second conflicting access. 第二个冲突
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}

/*
$ go run -race main.go
==================
WARNING: DATA RACE
Write at 0x00c0000c2450 by goroutine 7:
  runtime.mapassign_faststr()
      C:/Program Files/Go/src/runtime/map_faststr.go:202 +0x0
  main.main.func1()
      C:/workspace/go/go_demo/std/race/mypakg/main.go:9 +0x50

Previous write at 0x00c0000c2450 by main goroutine:
  runtime.mapassign_faststr()
      C:/Program Files/Go/src/runtime/map_faststr.go:202 +0x0
  main.main()
      C:/workspace/go/go_demo/std/race/mypakg/main.go:12 +0x131

Goroutine 7 (running) created at:
  main.main()
      C:/workspace/go/go_demo/std/race/mypakg/main.go:8 +0x114
==================
1 a
2 b
Found 1 data race(s)
exit status 66
*/
