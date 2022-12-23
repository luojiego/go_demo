package main

import (
	"net"
	"time"
)

// Concurrent reads and writes of the same map are not safe.
// Must use sync.Mutex go protect service
var service map[string]net.Addr

func RegisterService(name string, addr net.Addr) {
	service[name] = addr
}

func LookupService(name string) net.Addr {
	return service[name]
}

func main() {
	service = make(map[string]net.Addr)
	go func() {
		RegisterService("123", nil)
	}()

	time.Sleep(time.Second)
	LookupService("123")
}
