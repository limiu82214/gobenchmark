package map_slice

import (
	"math/rand"
	"testing"
)

// slice 在存取的速度上比 map 快
//l = 1e5
//BenchmarkPutAndGetSlice-10          1437            838788 ns/op               0 B/op          0 allocs/op
//BenchmarkPutAndGetMap-10             261           4873875 ns/op               0 B/op          0 allocs/op

func getRandomArray() []int {
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		arr[i] = i
	}
	return arr
}

const l = 1e5

/*
func BenchmarkPutAndGetArray(b *testing.B) {
	b.StopTimer()
	arr := getRandomArray()
	a := [l]int{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < l; j++ {
			// put
			idx := rand.Intn(l)
			a[idx] = arr[idx]

			// get
			if a[idx] != 0 {
			}
		}
	}

}
*/

func BenchmarkPutAndGetSlice(b *testing.B) {
	b.StopTimer()
	arr := getRandomArray()
	s := make([]int, l)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < l; j++ {
			// put
			idx := rand.Intn(l)
			s[idx] = arr[idx]

			// get
			if s[idx] != 0 {
			}
		}

	}
}

func BenchmarkPutAndGetMap(b *testing.B) {
	b.StopTimer()
	arr := getRandomArray()
	m := make(map[int]int, l)
	for i := 0; i < l; i++ {
		m[i] = 0
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < l; j++ {
			// put
			idx := rand.Intn(l)
			m[idx] = arr[idx]

			// get
			if _, ok := m[idx]; ok {
			}
		}
	}
}
