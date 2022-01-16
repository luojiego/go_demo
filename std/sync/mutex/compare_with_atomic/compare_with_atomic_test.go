package compare_with_atomic_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

var (
	n1 int64
)

func addSyncByAtomic(delta int64) int64 {
	return atomic.AddInt64(&n1, delta)
}

func readSyncByAtomic() int64 {
	return atomic.LoadInt64(&n1)
}

var (
	n2   int64
	rwmu sync.RWMutex
)

func addSyncByRWMutex(delta int64) {
	rwmu.Lock()
	defer rwmu.Unlock()
	n2 += delta
}

func readSyncByRWMutex() int64 {
	rwmu.RLock()
	defer rwmu.RUnlock()
	return n2
}

func BenchmarkAddSyncByAtomic(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			addSyncByAtomic(1)
		}
	})
}

func BenchmarkReadSyncByAtomic(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			readSyncByAtomic()
		}
	})
}

func BenchmarkAddSyncByMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			addSyncByRWMutex(1)
		}
	})
}

func BenchmarkReadSyncByMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			readSyncByRWMutex()
		}
	})
}

/*
Running tool: C:\Go\bin\go.exe test -benchmem -run=^$ -bench ^(BenchmarkAddSyncByAtomic|BenchmarkReadSyncByAtomic|BenchmarkAddSyncByMutex|BenchmarkReadSyncByMutex)$ compare_with_atomic -v

goos: windows
goarch: amd64
pkg: compare_with_atomic
cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
BenchmarkAddSyncByAtomic
BenchmarkAddSyncByAtomic-12
100000000	        17.97 ns/op	       0 B/op	       0 allocs/op
BenchmarkReadSyncByAtomic
BenchmarkReadSyncByAtomic-12
1000000000	         0.1463 ns/op	       0 B/op	       0 allocs/op
BenchmarkAddSyncByMutex
BenchmarkAddSyncByMutex-12
20964397	        57.11 ns/op	       0 B/op	       0 allocs/op
BenchmarkReadSyncByMutex
BenchmarkReadSyncByMutex-12
40696588	        30.52 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	compare_with_atomic	4.571s
*/
