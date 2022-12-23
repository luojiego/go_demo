package trace

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

// 获取 goroutine id 的方法
var goroutineSpace = []byte("goroutine ")

func curGoroutineID() uint64 {
	// bp := http2littleBuf.Get().(*[]byte)
	// defer http2littleBuf.Put(bp)
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	// Parse the 4707 out of "goroutine 4707 ["
	b = bytes.TrimPrefix(b, goroutineSpace)
	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		panic(fmt.Sprintf("No space found in %q", b))
	}
	b = b[:i]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
	}
	return n
}

func printTrace(id uint64, name, arrow string, indent int) {
	indents := ""
	for i := 0; i < indent; i++ {
		indents += "    "
	}
	fmt.Printf("g[%05d]:%s%s%s\n", id, indents, arrow, name)
}

var mu sync.Mutex
var m = make(map[uint64]int)

func Trace() func() {
	// 自动获取函数名称
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}
	fn := runtime.FuncForPC(pc)
	name := fn.Name()
	gid := curGoroutineID()

	mu.Lock()
	indents := m[gid]    // 获取当前 grid 对应的缩进层次
	m[gid] = indents + 1 // 缩进层次 +1 后存入 map
	mu.Unlock()
	// fmt.Printf("g[%05d]: enter: [%s]\n", gid, name)
	printTrace(gid, name, "->", indents+1)
	return func() {
		printTrace(gid, name, "<-", indents+1)
		// fmt.Printf("g[%05d]: exit: [%s]\n", gid, name)
	}
}
