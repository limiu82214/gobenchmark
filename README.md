# Golang 基準測試

Golang 的基準測試是一個用於測試代碼性能的工具，它通過測量某些操作的執行時間來幫助開發人員優化其程序。

## 運行基準測試

要運行基準測試，可以使用 `go test` 命令加上 `-bench` 參數。以下是一個範例命令:

```bash
go test -bench=.
```

在上面的命令中，`-bench` 參數後面的 `.` 表示運行所有基準測試。你也可以通過指定特定的測試名稱來運行單個基準測試。

運行基準測試的結果將會在終端機中顯示。以下是一個基準測試的範例結果：
```bash
goos: darwin
goarch: arm64
pkg: local
BenchmarkIsPalindrome-10 1000000000 0.7959 ns/op
PASS
ok local 0.985s
```

在這個範例中，`BenchmarkIsPalindrome` 表示執行的基準測試名稱。`-10` 表示測試併發數量。`1000000000` 表示總共運行了多少次測試。`0.7959 ns/op` 表示每次測試的平均耗時。

## 編寫基準測試

要編寫基準測試，需要在代碼中定義一個函數，並且使用 `Benchmark` 前綴來定義測試名稱。例如：

```go
func BenchmarkMyFunction(b *testing.B) {
    for i := 0; i < b.N; i++ {
        // 執行需要測試的代碼
    }
}
```
 
在上面的代碼中，BenchmarkMyFunction 是測試名稱。b 是一個 *testing.B 類型的指針，可以用來控制測試的執行次數。

注意，基準測試的結果可能會受到多種因素的影響，例如 CPU 時鐘速度、內存使用情況等。在進行基準測試時，應該儘可能消除這些因素的影響，以獲得更精確的結果。

* -cpuprofile=\$FILE：將 CPU 使用情況寫入指定的檔案 \$FILE 中，以便進行 CPU profiling。
* -memprofile=\$FILE：將記憶體使用情況寫入指定的檔案 \$FILE 中，以便進行記憶體 profiling。
* -memprofilerate=\$RATE：調整記憶體 profiling 的輸出量。預設為 512 KB，設定為 \$RATE 會讓輸出量變為預設值的 \$RATE 倍。
* -blockprofile=\$FILE：將 goroutine 的 blocking profile 寫入指定的檔案 \$FILE 中，以便進行 blocking profiling。
* -mutexprofile=\$FILE：將互斥鎖使用情況寫入指定的檔案 \$FILE 中，以便進行 mutex profiling。
