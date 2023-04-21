package redistest

// 要執行此測試必需要先安裝多台 Redis 伺服器，並設定好密碼、Port 等資訊
// personCount := 200, messageCount := 200, Count:   1,
// BenchmarkXReadMutipleMachine/1_redis-10         	       1	2160186583 ns/op	33963640 B/op	  686472 allocs/op
// BenchmarkXReadMutipleMachine/2_redis-10         	       1	1320368208 ns/op	33785360 B/op	  685109 allocs/op
// BenchmarkXReadMutipleMachine/3_redis-10         	       1	1014663291 ns/op	33066448 B/op	  676642 allocs/op
// BenchmarkXReadMutipleMachine/4_redis-10         	       2	 883806375 ns/op	32001480 B/op	  666604 allocs/op
// BenchmarkXReadMutipleMachine/5_redis-10         	       2	1157687792 ns/op	31837600 B/op	  663200 allocs/op
//
// personCount := 200, messageCount := 2, Count:   100, 觀查批次處理的效能
// BenchmarkXReadMutipleMachine/1_redis-10         	       1	1212091250 ns/op	78230144 B/op	 2074067 allocs/op
// BenchmarkXReadMutipleMachine/2_redis-10         	       2	 616941229 ns/op	76977936 B/op	 2068179 allocs/op
// BenchmarkXReadMutipleMachine/3_redis-10         	       3	 420165417 ns/op	76211162 B/op	 2047517 allocs/op
// BenchmarkXReadMutipleMachine/4_redis-10         	       3	 335048889 ns/op	75433306 B/op	 2026814 allocs/op
// BenchmarkXReadMutipleMachine/5_redis-10         	       4	 327825021 ns/op	75051914 B/op	 2016480 allocs/op
//
// personCount := 500, messageCount := 50, Count:   100, 觀查讀多送少的效能
// BenchmarkXReadMutipleMachine/1_redis-10         	       1	70940385500 ns/op	4813905696 B/op	129272847 allocs/op
// BenchmarkXReadMutipleMachine/2_redis-10         	       1	36016856875 ns/op	4812991960 B/op	129265650 allocs/op
// BenchmarkXReadMutipleMachine/3_redis-10         	       1	24493735166 ns/op	4793309736 B/op	128744380 allocs/op
// BenchmarkXReadMutipleMachine/4_redis-10         	       1	18749331083 ns/op	4773538840 B/op	128224357 allocs/op
// BenchmarkXReadMutipleMachine/5_redis-10         	       1	16200941834 ns/op	4763036088 B/op	127960635 allocs/op

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func benchmarkRedisClient(ctx context.Context, addr string, poolSize int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         addr,
		DialTimeout:  time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		PoolSize:     poolSize,
		Password:     "5ew4MTSL88GfEDgH",
	})
	return client
}

func BenchmarkXRead(b *testing.B) {
	ctx := context.Background()

	client := benchmarkRedisClient(ctx, ":63801", 100)
	defer client.Close()

	b.ResetTimer()

	var wg sync.WaitGroup

	personCount := 200
	messageCount := 2

	for j := 0; j < b.N; j++ {
		wg.Add(personCount)

		for i := 0; i < personCount; i++ {
			go readMessage(ctx, &wg, client, messageCount)
		}

		wg.Wait()
	}
}

func BenchmarkXReadMutipleMachine(b *testing.B) {
	ctx := context.Background()

	clientList := []*redis.Client{}
	clientList = append(clientList, benchmarkRedisClient(ctx, fmt.Sprintf(":%d", 63801), 100))
	defer clientList[0].Close()
	clientList = append(clientList, benchmarkRedisClient(ctx, fmt.Sprintf(":%d", 63802), 100))
	defer clientList[1].Close()
	clientList = append(clientList, benchmarkRedisClient(ctx, fmt.Sprintf(":%d", 63803), 100))
	defer clientList[2].Close()
	clientList = append(clientList, benchmarkRedisClient(ctx, fmt.Sprintf(":%d", 63804), 100))
	defer clientList[3].Close()
	clientList = append(clientList, benchmarkRedisClient(ctx, fmt.Sprintf(":%d", 63805), 100))
	defer clientList[4].Close()

	var wg sync.WaitGroup

	personCount := 200
	messageCount := 2

	for n := 1; n <= len(clientList); n++ {
		b.Run(fmt.Sprintf("%d redis", n), func(b *testing.B) {
			if personCount%n != 0 {
				personCount -= personCount % n
			}
			for j := 0; j < b.N; j++ {
				wg.Add(personCount)

				for k := 0; k < n; k++ {
					for i := 0; i < personCount/n; i++ {
						go readMessage(ctx, &wg, clientList[k], messageCount)
					}
				}

				wg.Wait()
			}
		})
	}
}

func readMessage(ctx context.Context, wg *sync.WaitGroup, client *redis.Client, times int) {
	defer wg.Done()
	var args = &redis.XReadArgs{
		Streams: []string{"guchat:mq", "0"},
		Count:   100,
		Block:   time.Duration(2) * time.Second,
	}
	for i := 0; i < times; i++ {
		_, err := client.XRead(ctx, args).Result()
		if err != nil {
			panic(err)
		}
	}
}
