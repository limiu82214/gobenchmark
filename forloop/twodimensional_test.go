package forloop

/**
 * 由於CPU特性 (CPU cache) 的關係，橫向的迴圈效能會比直向的迴圈效能好。
 * BenchmarkHorizontal-10                  24        47302252 ns/op    33333593 B/op      0 allocs/op
 * BenchmarkVertical-10                     3       404966486 ns/op   266668717 B/op      0 allocs/op
 */

import (
	"fmt"
	"testing"
)

func BenchmarkTwoDimensional(b *testing.B) {
	b.Run(fmt.Sprintf("Horizontal size-%d", 10000), func(b *testing.B) {
		arr := [10000][10000]int{}
		for i := 0; i < b.N; i++ {
			for x := 0; x < len(arr); x++ {
				for y := 0; y < len(arr); y++ {
					arr[x][y] = 1
				}
			}
		}
	})
	b.Run(fmt.Sprintf("Vertical size-%d", 10000), func(b *testing.B) {
		var arr [10000][10000]int
		for i := 0; i < b.N; i++ {
			for x := 0; x < len(arr); x++ {
				for y := 0; y < len(arr); y++ {
					arr[y][x] = 1
				}
			}
		}
	})
}
