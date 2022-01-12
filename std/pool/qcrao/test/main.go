package main

import (
	"fmt"
	"sync"
)

func main() {
	//没有提示 New 方法的时候 调用 Get 只会返回Nil
	//只有再 Put 了之后，才会正常的 Get 到对象
	var p sync.Pool

	g := p.Get()
	fmt.Println(g)
	//defer debug.SetGCPercent(debug.SetGCPercent(-1))
	//var p sync.Pool
	//if p.Get() != nil {
	//	log.Fatal("expected empty")
	//}
	//
	//// Make sure that the goroutine doesn't migrate to another P
	//// between Put and Get calls.
	////Runtime_procPin()
	//p.Put("a")
	//p.Put("b")
	//g := p.Get()
	//log.Printf("got %#v; want a\n", g)
	//
	//if g := p.Get(); g != "b" {
	//	log.Fatalf("got %#v; want b", g)
	//}
	//if g := p.Get(); g != nil {
	//	log.Fatalf("got %#v; want nil", g)
	//}
	////Runtime_procUnpin()
	//
	//// Put in a large number of objects so they spill into
	//// stealable space.
	//for i := 0; i < 100; i++ {
	//	p.Put("c")
	//}
	//// After one GC, the victim cache should keep them alive.
	//runtime.GC()
	//if g := p.Get(); g != "c" {
	//	log.Fatalf("got %#v; want c after GC", g)
	//}
	//// A second GC should drop the victim cache.
	//runtime.GC()
	//if g := p.Get(); g != nil {
	//	log.Fatalf("got %#v; want nil after second GC", g)
	//}
}
