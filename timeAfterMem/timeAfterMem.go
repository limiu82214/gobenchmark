package main

import (
	"fmt"
	"runtime"
	"time"
)

// time.After物件會在時間到達後，才有可能被GC回收
// 所以如果大量的請求，會導致記憶體使用量增加
// 這個問題應該會在go 1.23以後被修正

func consumer(outChan <-chan int) {
	for {
		select {
		//case <-time.After(60 * time.Second): //FIXME: 比較 Memory 的差異
		case <-time.After(10 * time.Millisecond): //FIXME: 比較 Memory 的差異
			fmt.Println("Timed check, no task received.")

		case val := <-outChan:
			_ = val
			//fmt.Printf("Processing value: %d\n", val)
		}
	}
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func main() {
	outChan := make(chan int)

	var memBefore runtime.MemStats
	runtime.ReadMemStats(&memBefore)
	fmt.Printf("Memory Usage Before sending data: Alloc = %v MiB, TotalAlloc = %v MiB, Sys = %v MiB, NumGC = %v\n",
		bToMb(memBefore.Alloc), bToMb(memBefore.TotalAlloc), bToMb(memBefore.Sys), memBefore.NumGC)

	go consumer(outChan)
	// 模拟生产者向通道发送数据
	for i := 0; i < 1000000; i++ {
		outChan <- i
	}

	var memAfter runtime.MemStats
	runtime.ReadMemStats(&memAfter)
	fmt.Printf("Memory Usage After sending data: Alloc = %v MiB, TotalAlloc = %v MiB, Sys = %v MiB, NumGC = %v\n",
		bToMb(memAfter.Alloc), bToMb(memAfter.TotalAlloc), bToMb(memAfter.Sys), memAfter.NumGC)

}
