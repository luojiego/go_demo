package main

import (
	"fmt"
	"unsafe"

	"github.com/Re-volution/sizestruct"
)

type test struct {
	a int32
	b string
	c map[string]int
	d int64 `ss:"-"`
}

type Header struct {
	length int32   // 除 length 之外，包数据长度
	flag   [2]byte // 包数据校验 'L' 'T'
	cmd    int32   // 命令字
	enCode byte    // 加密标识
	flow   int32   // 客户端 数据标识
	status int16   // 服务端 为 0 表示成功，非 0 则从全局错误码中找对应关系
}

func main() {
	var data = new(test)
	fmt.Println(sizestruct.SizeOf(data))
	fmt.Println(sizestruct.SizeTOf(data)) //Including type size 包括type的大小
	fmt.Println(unsafe.Sizeof(Header{}))
	fmt.Println(unsafe.Alignof(Header{}))
	fmt.Println(sizestruct.SizeOf(Header{}))
}
