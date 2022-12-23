package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/coreos/etcd/clientv3"
	recipe "github.com/coreos/etcd/contrib/recipes"
)

var (
	addr      = flag.String("addr", "http://127.0.0.1:2379", "etcd address")
	queueName = flag.String("name", "my-test-queue", "queue name")
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
	q := recipe.NewPriorityQueue(cli, *queueName)

	// 从命令行读取命令
	consoleScanner := bufio.NewScanner(os.Stdin)
	for consoleScanner.Scan() {
		actions := consoleScanner.Text()
		items := strings.Split(actions, " ")
		switch items[0] {
		case "push":
			if len(items) != 3 {
				fmt.Println("must set value and priority to push")
				continue
			}
			pr, err := strconv.Atoi(items[2]) // 读取优先级
			if err != nil {
				fmt.Println("must set uint16 as priority")
				continue
			}
			q.Enqueue(items[1], uint16(pr))
		case "pop":
			v, err := q.Dequeue()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(v)
		case "quit", "exit":
			return
		default:
			fmt.Println("unknown action")
		}
	}
}
