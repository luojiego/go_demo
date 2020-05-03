package main

import (
	"github.com/google/gops/agent"
	"log"
	"time"
)

func main() {
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Hour)
	//Tip of the Day
	//1. strings.ra -> //自动补全 ReplaceAll
	//2. Ctrl + Shift + J // Vim 的 J 指令
	//3. Shift + ⌘ + F7 在当前文件中高亮鼠标所处符号 ⌘ + G 向前寻找目标 Shift + ⌘ + G 向后寻找目标

}
