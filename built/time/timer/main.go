package main

import (
	"fmt"
	"time"
)

var (
	t *time.Timer
)

func main() {
	fmt.Println("hello world!", time.Now())
	t = time.NewTimer(-3 * time.Second)
	select {
	case <-t.C:
		fmt.Println("10s 到了", time.Now())
		t.Reset(5 * time.Second)
		select {
		case <-t.C:
			fmt.Println("10s 之后又 5s", time.Now())
		}
	}

	tm := time.Now()
	fmt.Println(tm.Year(), tm.Month(), tm.Day(), tm.Hour())
	waitTime := time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour()+8, 0, 0, 0, tm.Location()).Sub(tm)
	fmt.Println(waitTime)
	timer := time.NewTimer(waitTime)
	select {
	case <-timer.C:
		fmt.Println("timer out")
	}
}
