package endecode

/**
* 經過測試，發現在編解碼效能方面，使用JSON-iterator比起原生JSON有明顯的效能優勢，
* 而使用MapStruct + map[string]interface{} 的效能則較不理想。
* BenchmarkMapstructure/mapstructure-200-10                          32290             36161 ns/op           70400 B/op       1200 allocs/op
* BenchmarkMapstructure/mapstructure-map[string]-200-10                738           1600675 ns/op         2724797 B/op      34148 allocs/op
* BenchmarkMapstructure/jsoniterator-200-10                           2918            396048 ns/op          120102 B/op       5400 allocs/op
* BenchmarkMapstructure/json-200-10                                   1693            699257 ns/op          163281 B/op       4200 allocs/op
*
* BenchmarkMapstructure/mapstructure-20000-10                          334           3592924 ns/op         7040018 B/op     120000 allocs/op
* BenchmarkMapstructure/mapstructure-map[string]-20000-10                7         160034899 ns/op        272468707 B/op   3414765 allocs/op
* BenchmarkMapstructure/jsoniterator-20000-10                           28          39748098 ns/op        12010486 B/op     540022 allocs/op
* BenchmarkMapstructure/json-20000-10                                   16          70060857 ns/op        16328366 B/op     420030 allocs/op
 */

import (
	"encoding/json"
	"fmt"
	"testing"

	jsoniterator "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"
)

type testStruct struct {
	A1 string
	A2 string
	A3 string
	A4 string
	A5 string
	B1 bool
	B2 bool
	B3 bool
	B4 bool
	B5 bool
	C1 int
	C2 int
	C4 int
	C5 int
	D1 float64
	D2 float64
	D3 float64
	D4 float64
	D5 float64
}

func BenchmarkDeEncode(b *testing.B) { //nolint:funlen,gocognit //for different method
	t := testStruct{
		A1: "value1",
		A2: "value2",
		A3: "value3",
		A4: "value4",
		A5: "value5",
		B1: true,
		B2: false,
		B3: true,
		B4: false,
		B5: true,
		C1: 10,
		C2: 20,
		C4: 40,
		C5: 50,
		D1: 1.23,
		D2: 4.56,
		D3: 7.89,
		D4: 0.12,
		D5: 3.45,
	}

	for _, count := range []int{200, 20000} {
		b.Run(fmt.Sprintf("mapstructure-%d", count), func(b *testing.B) {
			for j := 0; j < count; j++ {
				for i := 0; i < b.N; i++ {
					var v interface{}

					err := mapstructure.Decode(v, &t)
					if err != nil {
						panic(err)
					}

					ans := testStruct{}

					err = mapstructure.Decode(ans, &v)
					if err != nil {
						panic(err)
					}
				}
			}
		})
		b.Run(fmt.Sprintf("mapstructure-map[string]-%d", count), func(b *testing.B) {
			for j := 0; j < count; j++ {
				for i := 0; i < b.N; i++ {
					var v map[string]interface{}

					err := mapstructure.Decode(v, &t)
					if err != nil {
						panic(err)
					}

					ans := testStruct{}

					err = mapstructure.Decode(ans, &v)
					if err != nil {
						panic(err)
					}
				}
			}
		})
		b.Run(fmt.Sprintf("jsoniterator-%d", count), func(b *testing.B) {
			for j := 0; j < count; j++ {
				for i := 0; i < b.N; i++ {
					b, err := jsoniterator.Marshal(t)
					if err != nil {
						panic(err)
					}

					ans := testStruct{}

					err = jsoniterator.Unmarshal(b, &ans)
					if err != nil {
						panic(err)
					}
				}
			}
		})
		b.Run(fmt.Sprintf("json-%d", count), func(b *testing.B) {
			for j := 0; j < count; j++ {
				for i := 0; i < b.N; i++ {
					b, err := json.Marshal(t)
					if err != nil {
						panic(err)
					}

					ans := testStruct{}

					err = json.Unmarshal(b, &ans)
					if err != nil {
						panic(err)
					}
				}
			}
		})
	}
}
