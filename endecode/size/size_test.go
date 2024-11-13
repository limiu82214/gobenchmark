package size

import (
	"encoding/json"
	"fmt"
	"github.com/limiu82214/gobenchmark/endecode/size/myproto"
	"google.golang.org/protobuf/proto"
	"testing"
)

// 使用json和proto進行編解碼，並比較其效能
// pkg: github.com/limiu82214/gobenchmark/endecode/size
// BenchmarkDeEncode/json-200-10               1561            692813 ns/op          156885 B/op       2400 allocs/op
// BenchmarkDeEncode/proto-200-10             12800             93786 ns/op           62400 B/op       1400 allocs/op
// BenchmarkDeEncode/json-20000-10               18          63604338 ns/op        15688506 B/op     240036 allocs/op
// BenchmarkDeEncode/proto-20000-10             127           9396231 ns/op         6240062 B/op     140000 allocs/op
// json size: 205
// proto size: 83
// PASS
// ok      github.com/limiu82214/gobenchmark/endecode/size 6.918s

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
	tProto := myproto.TestStruct{
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
		b.Run(fmt.Sprintf("proto-%d", count), func(b *testing.B) {
			for j := 0; j < count; j++ {
				for i := 0; i < b.N; i++ {
					b, err := proto.Marshal(&tProto)
					if err != nil {
						panic(err)
					}

					ans := myproto.TestStruct{}

					err = proto.Unmarshal(b, &ans)
					if err != nil {
						panic(err)
					}
				}
			}
		})
	}

	// json Size
	fJ := func() {
		b, err := json.Marshal(t)
		if err != nil {
			panic(err)
		}

		ans := testStruct{}

		err = json.Unmarshal(b, &ans)
		if err != nil {
			panic(err)
		}
		fmt.Println("json size:", len(b))
	}
	fJ()

	// proto Size
	fP := func() {
		b, err := proto.Marshal(&tProto)
		if err != nil {
			panic(err)
		}

		ans := myproto.TestStruct{}

		err = proto.Unmarshal(b, &ans)
		if err != nil {
			panic(err)
		}
		fmt.Println("proto size:", len(b))
	}
	fP()
}
