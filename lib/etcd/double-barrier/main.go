package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	recipe "github.com/coreos/etcd/contrib/recipes"
)

var (
	addr        = flag.String("addr", "http://127.0.0.1:2379", "etcd address")
	barrierName = flag.String("name", "my-test-queue", "barrier name")
	count       = flag.Int("c", 2, "")
)

func main() {
	flag.Parse()
	endpoints := strings.Split(*addr, ",")

	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// 创建 session
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()

	// 创建/获取队列
	b := recipe.NewDoubleBarrier(s1, *barrierName, *count)

	// 从命令行读取命令
	consoleScanner := bufio.NewScanner(os.Stdin)
	for consoleScanner.Scan() {
		actions := consoleScanner.Text()
		items := strings.Split(actions, " ")
		switch items[0] {
		case "enter": // 当调用者调用 Enter 时，会被阻塞住，直到一共有 count 个节点调用 Enter 时，这 count 个阻塞的节点才能继续执行
			b.Enter()
			fmt.Println("enter")
		case "leave":
			b.Leave() // 打开这个栅栏
			fmt.Println("leave")
		case "quit", "exit":
			return
		default:
			fmt.Println("unknown action")
		}
	}
}
