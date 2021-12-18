package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"sync"
	"t1/testproto"
)

func main() {
	t := &testproto.FlyObjectCate{
		FlyCateMap: map[uint32]uint32{1: 2, 3: 3},
		Name:       "luojie",
	}
	res, _ := json.Marshal(t)
	fmt.Println(string(res))

	var m sync.Map
	var a, b, c, d, e uint32 = 1, 2, 3, 0, 100
	m.LoadOrStore(a, "a")
	m.LoadOrStore(b, "b")
	m.LoadOrStore(c, "c")
	m.LoadOrStore(d, "d")
	m.LoadOrStore(e, "e")
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("key: %v value: %v\n", key, value)
		return true
	})

	var keys []int
	m.Range(func(key, _ interface{}) bool {
		keys = append(keys, int(key.(uint32)))
		return true
	})
	fmt.Println(keys)
	sort.Ints(keys)
	fmt.Println(keys)
}
