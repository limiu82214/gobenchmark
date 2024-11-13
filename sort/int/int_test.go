package int

import (
	"cmp"
	"math"
	"math/rand"
	"slices"
	"sort"
	"testing"
)

const l = 1000000

func getRandomArray() []int {
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		arr[i] = rand.Intn(math.MaxInt) - (math.MaxInt / 2)
	}
	return arr
}

func BenchmarkSortSlice(b *testing.B) {
	b.StopTimer()
	arr := getRandomArray()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		//arr[rand.Intn(l)] = rand.Intn(math.MaxInt)
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
	}
}

func BenchmarkSortInts(b *testing.B) {
	b.StopTimer()
	arr := getRandomArray()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		//arr[rand.Intn(l)] = rand.Intn(math.MaxInt)
		sort.Ints(arr)
	}
}

func SliceDiff[T comparable](a, b []T) []T {
	m := make(map[T]struct{}, len(b))
	for _, x := range b {
		m[x] = struct{}{}
	}
	var diff []T
	for _, x := range a {
		if _, ok := m[x]; !ok {
			diff = append(diff, x)
		}
	}
	return diff
}

func BenchmarkSlicesSortCompareFunc(b *testing.B) {
	b.StopTimer()
	arr := getRandomArray()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		//arr[rand.Intn(l)] = rand.Intn(math.MaxInt)
		slices.SortFunc(arr, func(a, b int) int {
			return cmp.Compare(a, b)
		})
	}
}
func BenchmarkSlicesSortSubFunc(b *testing.B) {
	b.StopTimer()
	arr := getRandomArray()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		//arr[rand.Intn(l)] = rand.Intn(math.MaxInt)
		slices.SortFunc(arr, func(a, b int) int {
			return a - b
		})
	}
}

func BenchmarkSlicesSort(b *testing.B) {
	b.StopTimer()
	arr := getRandomArray()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		//arr[rand.Intn(l)] = rand.Intn(math.MaxInt)
		slices.Sort(arr)
	}
}

func BenchmarkRadixSort(b *testing.B) {
	b.StopTimer()
	arr := getRandomArray()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		//arr[rand.Intn(l)] = rand.Intn(math.MaxInt)
		radixSort(arr)
	}
}

func radixSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// 找出數組中的最大值和最小值，以便確定數值範圍
	minVal, maxVal := arr[0], arr[0]
	for _, v := range arr {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}

	// 將數組中的所有值轉換為非負整數，方便處理
	offset := 0
	if minVal < 0 {
		offset = -minVal
		for i := 0; i < n; i++ {
			arr[i] += offset
		}
		maxVal += offset
	}

	// 開始基數排序
	tmp := make([]int, n)
	radix := 256 // 基數，對應一個位元組的可能值數量
	for shift := 0; (maxVal >> shift) > 0; shift += 8 {
		count := [256]int{}

		// 計數每個位元組值的出現次數
		for i := 0; i < n; i++ {
			b := (arr[i] >> shift) & 0xFF
			count[b]++
		}

		// 計算累積和
		sum := 0
		for i := 0; i < radix; i++ {
			c := count[i]
			count[i] = sum
			sum += c
		}

		// 根據位元組值進行排序
		for i := 0; i < n; i++ {
			b := (arr[i] >> shift) & 0xFF
			tmp[count[b]] = arr[i]
			count[b]++
		}

		// 將排序結果複製回原始陣列
		copy(arr, tmp)
	}

	// 如果有偏移，將數值還原
	if offset > 0 {
		for i := 0; i < n; i++ {
			arr[i] -= offset
		}
	}
}
