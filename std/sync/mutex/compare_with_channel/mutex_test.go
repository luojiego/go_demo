package mutex_compare_with_chan

import (
	"sync"
	"testing"
)

var (
	cs = 0
	mu sync.Mutex
	c  = make(chan struct{}, 1)
)

func criticalSectionSyncByMutex() {
	mu.Lock()
	defer mu.Unlock()
	cs++
}

func criticalSectionSyncByChan() {
	c <- struct{}{}
	cs++
	<-c
}

func BenchmarkCriticalSectionSyncByMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		criticalSectionSyncByMutex()
	}
}

func BenchmarkCriticalSectionSyncByMutexInParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			criticalSectionSyncByMutex()
		}
	})
}

func BenchmarkCriticalSectionSyncByChan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		criticalSectionSyncByChan()
	}
}

func BenchmarkCriticalSectionSyncByChanInParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			criticalSectionSyncByChan()
		}
	})
}

/*
goos: windows
goarch: amd64
pkg: compare_with_channel
cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
BenchmarkCriticalSectionSyncByMutex
BenchmarkCriticalSectionSyncByMutex-12
88626946	        13.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkCriticalSectionSyncByMutexInParallel
BenchmarkCriticalSectionSyncByMutexInParallel-12
26813306	        44.46 ns/op	       0 B/op	       0 allocs/op
BenchmarkCriticalSectionSyncByChan
BenchmarkCriticalSectionSyncByChan-12
39299808	        36.70 ns/op	       0 B/op	       0 allocs/op
BenchmarkCriticalSectionSyncByChanInParallel
BenchmarkCriticalSectionSyncByChanInParallel-12
 5991354	       200.7 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	compare_with_channel	6.331s
*/
