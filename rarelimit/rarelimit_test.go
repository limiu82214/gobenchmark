package rarelimit_test

import (
	"fmt"
	"golang.org/x/time/rate"
	"testing"
	"time"
)

var limiter = rate.NewLimiter(rate.Limit(1), 5) // 每秒 1 次请求，最多突发 5 次

// 限制函数调用频率的包装器
func rateLimitedFoo() {
	if !limiter.Allow() {
		fmt.Println("rate limit exceeded")
		return
	}

	// 被限制的
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!foo called")
}

func main() {
}

// simple
func BenchmarkRateLimitedFooSimple(b *testing.B) {
	// 手动测试函数调用

	for i := 0; i < 30; i++ {
		rateLimitedFoo()
		time.Sleep(333 * time.Millisecond) // 每次调用间隔 200 毫秒
	}
}

//// 基准测试
//func BenchmarkRateLimitedFoo(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		rateLimitedFoo()
//	}
//}
//
//func BenchmarkRateLimitedFooParallel(b *testing.B) {
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next() {
//			rateLimitedFoo()
//		}
//	})
//}
