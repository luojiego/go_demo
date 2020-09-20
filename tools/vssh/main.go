package main

import (
	"context"
	"fmt"
	"github.com/yahoo/vssh"
	"log"
	"time"
)

func main() {
	vs := vssh.New().Start()
	config, _ := vssh.GetConfigPEM("vssh", "mypem.pem")
	vs.AddClient("35.220.193.53:22", config, vssh.SetMaxSessions(4))
	vs.Wait()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// cmd:= "ping -c 4 192.168.55.10"
	cmd := "ping -c 10.170.0.3"
	timeout, _ := time.ParseDuration("6s")
	respChan := vs.Run(ctx, cmd, timeout)

	resp := <-respChan
	if err := resp.Err(); err != nil {
		log.Fatal(err)
	}

	stream := resp.GetStream()
	defer stream.Close()

	for stream.ScanStdout() {
		txt := stream.TextStdout()
		fmt.Println(txt)
	}
}
