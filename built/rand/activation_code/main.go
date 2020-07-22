package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//日本包需求 需要生成激活码

//十位激活码 一次最多生成50000个

//前4位用做序号 为001-999 base64之后的结果
//后6位用随机生成激活码

var (
	s = []rune("23456789abcdefghijkmnpqrstuvwxyz")
	m = make(map[string]bool, 0)
)

func main() {
	fmt.Println(len(s))

	rand.Seed(time.Now().UnixNano())

	count := 0

	//encoder := base64.NewEncoding()
	//r := encoder.EncodeToString([]byte("001"))
	//fmt.Println(r)
	batch := fmt.Sprintf("%04d", 1)
	for i := 0; i < 50000; i++ {
		result := RandStringRunes(batch, 8)
		fmt.Println(result)
		if _, ok := m[result]; ok {
			//m[result]
			count++
		} else {
			m[result] = true
		}
	}
	fmt.Println("count: ", count)
}

func RandStringRunes(batch string, n int) string {
	batch = strings.ReplaceAll(batch, "0", "a")
	batch = strings.ReplaceAll(batch, "1", "b")
	b := make([]rune, n)
	for i := range b {
		b[i] = s[rand.Intn(len(s))]
	}
	return batch + string(b)
}
