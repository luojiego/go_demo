package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	w := &Watchdog{}

	go func() {
		w.Start()
	}()

	w.KeepAlive()

	time.Sleep(10 * time.Second)
}

type Watchdog struct {
	last int64
}

func (w *Watchdog) KeepAlive() {
	w.last = time.Now().UnixNano() // First conflicting access.
	// right process
	// atomic.StoreInt64(&w.last, time.Now().UnixNano())
}

func (w *Watchdog) Start() {
	go func() {
		for {
			time.Sleep(time.Second)
			// second conflicting access

			// if atomic.LoadInt64(&w.last) < time.Now().Add(-10*time.Second).UnixNano() {
			if w.last < time.Now().Add(-10*time.Second).UnixNano() {
				fmt.Println("No Keepalives for 10 seconds. Dying.")
				os.Exit(1)
			}
		}
	}()
}
