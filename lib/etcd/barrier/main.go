package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/coreos/etcd/clientv3"
	recipe "github.com/coreos/etcd/contrib/recipes"
)

var (
	addr        = flag.String("addr", "http://127.0.0.1:2379", "etcd address")
	barrierName = flag.String("name", "my-test-queue", "barrier name")
)

func main() {
	flag.Parse()
	endpoints := strings.Split(*addr, ",")

	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// 创建/获取队列
	b := recipe.NewBarrier(cli, *barrierName)

	// 从命令行读取命令
	consoleScanner := bufio.NewScanner(os.Stdin)
	for consoleScanner.Scan() {
		actions := consoleScanner.Text()
		items := strings.Split(actions, " ")
		switch items[0] {
		case "hold":
			b.Hold()
			fmt.Println("hold")
		case "release":
			b.Release() // 打开这个栅栏
			fmt.Println("released")
		case "wait":
			b.Wait()
			fmt.Println("after wait")
		case "quit", "exit":
			return
		default:
			fmt.Println("unknown action")
		}
	}
}
