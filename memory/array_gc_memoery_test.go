package memory

import (
	"runtime"
	"testing"
)

func TestPrefixMemGC(t *testing.T) {
	var memStatsBefore, memStatsAfter, memStatsAfterGC runtime.MemStats

	runtime.ReadMemStats(&memStatsBefore)
	arr := make([]int, 100000000000)
	//t.Log("ptr:", &arr[0])
	arr = arr[99999999999:]
	//t.Log("ptr:", &arr[0])
	runtime.ReadMemStats(&memStatsAfter)
	runtime.GC()
	runtime.ReadMemStats(&memStatsAfterGC)

	//t.Log("ptr:", &arr[0]) // 注解開這行，他就不會被回收
	t.Log("cap:", cap(arr), "len:", len(arr))
	t.Logf("Before: %d bytes", memStatsBefore.HeapAlloc)
	t.Logf("After: %d bytes", memStatsAfter.HeapAlloc)
	t.Logf("After-GC: %d bytes", memStatsAfterGC.HeapAlloc)
}

func TestSuffixMemGC(t *testing.T) {
	var memStatsBefore, memStatsAfter, memStatsAfterGC runtime.MemStats

	runtime.ReadMemStats(&memStatsBefore)
	arr := make([]int, 100000000000)
	//t.Log("ptr:", &arr[0])
	arr = arr[:1]
	//t.Log("ptr:", &arr[0])
	runtime.ReadMemStats(&memStatsAfter)
	runtime.GC()
	runtime.ReadMemStats(&memStatsAfterGC)

	//t.Log("ptr:", &arr[0]) // 注解開這行，他就不會被回收
	t.Log("cap:", cap(arr), "len:", len(arr))
	t.Logf("Before: %d bytes", memStatsBefore.HeapAlloc)
	t.Logf("After: %d bytes", memStatsAfter.HeapAlloc)
	t.Logf("After-GC: %d bytes", memStatsAfterGC.HeapAlloc)
}
