package switchTest

import (
	"fmt"
	"sync"
	"testing"
)

// 總共做totla次運算，分成nums個goroutine去做，觀察goroutine切換的效能
func BenchmarkSwitching(b *testing.B) {
	total := int(1e9)
	nums := []int{1e3, 1e6, 5e6}

	for n := 0; n < len(nums); n++ {
		b.Run(fmt.Sprintf("use %d goroutines", nums[n]), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var wg sync.WaitGroup
				wg.Add(nums[n])
				for i := 0; i < nums[n]; i++ {
					go func() {
						for j := 0; j < total/nums[n]; j++ {
							_ = j * j
						}
						wg.Done()
					}()
				}
				wg.Wait()
			}
		})
	}
}
