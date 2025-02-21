package main

import (
	"runtime"
	"testing"
)

const l = 1e9

func BenchmarkPrefixMemGC(b *testing.B) {
	var totalBefore, totalAfter, totalAfterGC uint64

	for i := 0; i < b.N; i++ {
		var memStatsBefore, memStatsAfter, memStatsAfterGC runtime.MemStats

		runtime.ReadMemStats(&memStatsBefore)
		arr := make([]int, l)
		arr = arr[l-1:] // 去除前面的元素
		runtime.ReadMemStats(&memStatsAfter)
		runtime.GC() // 手動觸發 GC
		runtime.ReadMemStats(&memStatsAfterGC)

		totalBefore += memStatsBefore.HeapAlloc
		totalAfter += memStatsAfter.HeapAlloc
		totalAfterGC += memStatsAfterGC.HeapAlloc
	}

	avgBefore := totalBefore / uint64(b.N)
	avgAfter := totalAfter / uint64(b.N)
	avgAfterGC := totalAfterGC / uint64(b.N)

	b.Log("Benchmark Prefix Mem GC")
	b.Logf("Average cap: %d, len: %d", cap(make([]int, l)[l-1:]), len(make([]int, l)[l-1:]))
	b.Logf("Average Before: %d bytes", avgBefore)
	b.Logf("Average After: %d bytes", avgAfter)
	b.Logf("Average After-GC: %d bytes", avgAfterGC)
}

func BenchmarkSuffixMemGC(b *testing.B) {
	var totalBefore, totalAfter, totalAfterGC uint64

	for i := 0; i < b.N; i++ {
		var memStatsBefore, memStatsAfter, memStatsAfterGC runtime.MemStats

		runtime.ReadMemStats(&memStatsBefore)
		arr := make([]int, l)
		arr = arr[:1] // 去除後面的元素
		runtime.ReadMemStats(&memStatsAfter)
		runtime.GC() // 手動觸發 GC
		runtime.ReadMemStats(&memStatsAfterGC)

		totalBefore += memStatsBefore.HeapAlloc
		totalAfter += memStatsAfter.HeapAlloc
		totalAfterGC += memStatsAfterGC.HeapAlloc
	}

	avgBefore := totalBefore / uint64(b.N)
	avgAfter := totalAfter / uint64(b.N)
	avgAfterGC := totalAfterGC / uint64(b.N)

	b.Log("Benchmark Suffix Mem GC")
	b.Logf("Average cap: %d, len: %d", cap(make([]int, l)[:1]), len(make([]int, l)[:1]))
	b.Logf("Average Before: %d bytes", avgBefore)
	b.Logf("Average After: %d bytes", avgAfter)
	b.Logf("Average After-GC: %d bytes", avgAfterGC)
}
