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
 go test -cpu=2,4,8,16 -benchmem -bench .
goos: windows
goarch: amd64
pkg: compare_with_atomic
cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
BenchmarkAddSyncByAtomic-2              92093156                15.19 ns/op            0 B/op          0 allocs/op
BenchmarkAddSyncByAtomic-4              84246589                13.85 ns/op            0 B/op          0 allocs/op
BenchmarkAddSyncByAtomic-8              100000000               14.86 ns/op            0 B/op          0 allocs/op
BenchmarkAddSyncByAtomic-16             100000000               14.91 ns/op            0 B/op          0 allocs/op
BenchmarkReadSyncByAtomic-2             1000000000               0.6405 ns/op          0 B/op          0 allocs/op
BenchmarkReadSyncByAtomic-4             1000000000               0.3295 ns/op          0 B/op          0 allocs/op
BenchmarkReadSyncByAtomic-8             1000000000               0.1835 ns/op          0 B/op          0 allocs/op
BenchmarkReadSyncByAtomic-16            1000000000               0.1315 ns/op          0 B/op          0 allocs/op
BenchmarkAddSyncByMutex-2               40963802                29.28 ns/op            0 B/op          0 allocs/op
BenchmarkAddSyncByMutex-4               25370192                46.93 ns/op            0 B/op          0 allocs/op
BenchmarkAddSyncByMutex-8               21890478                55.55 ns/op            0 B/op          0 allocs/op
BenchmarkAddSyncByMutex-16              20860033                57.32 ns/op            0 B/op          0 allocs/op
BenchmarkReadSyncByMutex-2              38162382                36.28 ns/op            0 B/op          0 allocs/op
BenchmarkReadSyncByMutex-4              36196146                36.72 ns/op            0 B/op          0 allocs/op
BenchmarkReadSyncByMutex-8              35981361                36.47 ns/op            0 B/op          0 allocs/op
BenchmarkReadSyncByMutex-16             30544843                35.93 ns/op            0 B/op          0 allocs/op
PASS
ok      compare_with_atomic     23.266s
*/
