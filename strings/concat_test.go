package strings

/**
 * 使用 + 號進行字串連接的效能比使用 fmt.Sprintf 要好得多。
 * BenchmarkStringConcat-10           36440635        32.65 ns/op      16 B/op         1 allocs/op
 * BenchmarkStringConcatSprintf-10    10311322        106.7 ns/op      32 B/op         3 allocs/op
 */

import (
	"fmt"
	"testing"
)

func BenchmarkStringConcat(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = ""
		s += ""
		s += "id: " + "1"
		s += "name: " + "jack" //nolint:ineffassign // test
	}
}

func BenchmarkStringConcatSprintf(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = ""
		s += fmt.Sprintf("id: %v", "1")
		s += fmt.Sprintf("name: %v", "jack") //nolint:ineffassign // test
	}
}
