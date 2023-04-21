package channell

/**
 * 調整gorountineCount的數量，可以看到效能的差異，
 * 當gorountineCount太多時，會造成gorountine的切換與GC回收，導致效能下降。
 */

import (
	"sync"
	"testing"
	"time"
)

func BenchmarkTooManyWrite(b *testing.B) {
	var ch = make(chan int, 128)
	defer close(ch)

	wg := sync.WaitGroup{}
	gorountineCount := 10000 * 50

	for i := 0; i < gorountineCount; i++ {
		wg.Add(1)

		go func() {
			ch <- 12313

			wg.Done()
		}()
	}

	go func() {
		for _ = range ch {
		}
	}()
	wg.Wait()
}

func BenchmarkTooManyWriteStruct(b *testing.B) { //nolint:funlen // because test struct is too long
	gorountineCount := 10000 * 20

	type ExampleStruct struct {
		User struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"user"`
		Cart struct {
			Items []struct {
				Product struct {
					ID          string  `json:"id"`
					Name        string  `json:"name"`
					Description string  `json:"description"`
					Price       float64 `json:"price"`
				} `json:"product"`
				Quantity int `json:"quantity"`
			} `json:"items"`
			Total float64 `json:"total"`
		} `json:"cart"`
		Shipping struct {
			Method  string `json:"method"`
			Address struct {
				Street string `json:"street"`
				City   string `json:"city"`
				State  string `json:"state"`
				Zip    string `json:"zip"`
			} `json:"address"`
			Fee float64 `json:"fee"`
		} `json:"shipping"`
		CreateAt  time.Time `json:"create_at"`
		RequestID string    `json:"request_id"`
	}

	var ch = make(chan *ExampleStruct, 128)
	defer close(ch)

	wg := sync.WaitGroup{}

	for i := 0; i < gorountineCount; i++ {
		wg.Add(1)

		go func() {
			obj := &ExampleStruct{
				User: struct {
					Name  string "json:\"name\""
					Email string "json:\"email\""
				}{
					Name:  "jack",
					Email: "qwe@qwe.com",
				},
				Cart: struct {
					Items []struct {
						Product struct {
							ID          string  "json:\"id\""
							Name        string  "json:\"name\""
							Description string  "json:\"description\""
							Price       float64 "json:\"price\""
						} "json:\"product\""
						Quantity int "json:\"quantity\""
					} "json:\"items\""
					Total float64 "json:\"total\""
				}{
					Items: []struct {
						Product struct {
							ID          string  "json:\"id\""
							Name        string  "json:\"name\""
							Description string  "json:\"description\""
							Price       float64 "json:\"price\""
						} "json:\"product\""
						Quantity int "json:\"quantity\""
					}{},
					Total: 0.0,
				},
				Shipping: struct {
					Method  string "json:\"method\""
					Address struct {
						Street string "json:\"street\""
						City   string "json:\"city\""
						State  string "json:\"state\""
						Zip    string "json:\"zip\""
					} "json:\"address\""
					Fee float64 "json:\"fee\""
				}{
					Method: "",
					Address: struct {
						Street string "json:\"street\""
						City   string "json:\"city\""
						State  string "json:\"state\""
						Zip    string "json:\"zip\""
					}{
						Street: "",
						City:   "",
						State:  "",
						Zip:    "",
					},
					Fee: 0.0,
				},
				CreateAt:  time.Time{},
				RequestID: "",
			}
			ch <- obj

			wg.Done()
		}()
	}

	go func() {
		for _ = range ch {
		}
	}()
	wg.Wait()
}
