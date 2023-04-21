package mapp

/**
 * 這個 benchmark 測試了兩種不同的 map 實作方法：一種是不使用指標，另一種是使用指標。
 * 由結果可知，使用指標的方法比不使用指標的方法更慢，因為前者在每次迭代時都需要額外的記憶體分配和指標操作。在這兩種方法中，不使用指標的方法比使用指標的方法更有效率
 * * 會造成GC壓力
 * BenchmarkMapWithoutPtrs-10     203191    4939 ns/op    5221 B/op     14 allocs/op
 * BenchmarkMapWithPtrs-10        194050    6159 ns/op    5906 B/op    107 allocs/op
 */

import "testing"

func BenchmarkMapWithoutPtrs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var m = make(map[int]int)

		for i := 0; i < 100; i++ {
			m[i] = i
		}
	}
}

func BenchmarkMapWithPtrs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var m = make(map[int]*int)

		for i := 0; i < 100; i++ {
			var v = i
			m[i] = &v
		}
	}
}
