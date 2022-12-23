package main

import (
	"encoding/binary"
	"fmt"

	// "unsafe"
	"github.com/Re-volution/sizestruct"

	unsafe "unsafe"
)

type Header struct {
	length int32   // 除 length 之外，包数据长度
	flag   [2]byte // 包数据校验 (HeaderFlag)
	cmd    int32   // 命令字
	enCode byte    // 加密标识
	flow   int32   // 客户端 数据标识
	status int16   // 服务端 为 0 表示成功，非 0 则从全局错误码中找对应关系
}

func main() {
	fmt.Println("Hello, playground")
	h := Header{}
	fmt.Println(unsafe.Sizeof(h))
	fmt.Println(sizestruct.SizeOf(Header{}))

	// r := reflect.ValueOf(h)
	s := binary.Size(h)
	fmt.Println(s)
}
