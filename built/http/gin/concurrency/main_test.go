package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"sync"
	"testing"
)

var (
	mu  = sync.Mutex{}
	arr = make([]int, 0)
)

func getResult(index int, wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := http.Get("http://localhost:7001/")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	mu.Lock()
	defer mu.Unlock()
	n, _ := strconv.Atoi(string(result))
	arr = append(arr, n)

	//fmt.Printf("index: %d result: %d\n", index, n)
}

func TestConcurrency(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go getResult(i, &wg)
	}

	wg.Wait()
	println(len(arr))
	sort.Sort(sort.IntSlice(arr))
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1]+1 {
			panic(fmt.Errorf("data error arr[%d]=%d, arr[%d]=%d arr:%+v",
				i-1, arr[i-1], i, arr[i], arr[:(i+10)]))
		}
	}
}
