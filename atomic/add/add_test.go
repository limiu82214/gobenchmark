package add

import (
	"sync"
	"sync/atomic"
	"testing"
)

// 在大量併發的情況下，count++ > atomic > mutex
// 但直接使用count++需要考慮data race的問題
//var numGoroutines = 500
//var numIncrements = 100000
//BenchmarkAtomicAdd-10                  1        3271295458 ns/op          257896 B/op       1018 allocs/op
//BenchmarkNormalAdd-10                 88          15723227 ns/op           12138 B/op        500 allocs/op
//BenchmarkMutexAdd-10                   1        5855801124 ns/op          136128 B/op       1721 allocs/op
//PASS
//ok      github.com/limiu82214/gobenchmark/atomic/add    10.799s

var numGoroutines = 500
var numIncrements = 100000

func BenchmarkAtomicAdd(b *testing.B) {

	for t := 0; t < b.N; t++ {
		b.StopTimer()
		var count uint64
		var wg sync.WaitGroup
		wg.Add(numGoroutines)

		b.StartTimer()
		for i := 0; i < numGoroutines; i++ {
			go func() {
				defer wg.Done()
				for j := 0; j < numIncrements; j++ {
					atomic.AddUint64(&count, 1)
				}
			}()
		}
		wg.Wait()
		b.StopTimer()
	}
}
func BenchmarkNormalAdd(b *testing.B) {

	for t := 0; t < b.N; t++ {
		b.StopTimer()
		var count uint64
		var wg sync.WaitGroup
		wg.Add(numGoroutines)

		b.StartTimer()
		for i := 0; i < numGoroutines; i++ {
			go func() {
				defer wg.Done()
				for j := 0; j < numIncrements; j++ {
					count++
				}
			}()
		}
		wg.Wait()
		b.StopTimer()
	}
}

func BenchmarkMutexAdd(b *testing.B) {

	for t := 0; t < b.N; t++ {
		b.StopTimer()
		var count uint64
		var wg sync.WaitGroup
		var mu sync.Mutex
		wg.Add(numGoroutines)

		b.StartTimer()
		for i := 0; i < numGoroutines; i++ {
			go func() {
				defer wg.Done()
				for j := 0; j < numIncrements; j++ {
					mu.Lock()
					count++
					mu.Unlock()
				}
			}()
		}
		wg.Wait()
		b.StopTimer()
	}
}
