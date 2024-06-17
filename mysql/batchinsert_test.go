package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 10000筆 測試結果
//BenchmarkInsertBatch1-10               1        386506066083 ns/op       4722664 B/op     149773 allocs/op
//BenchmarkInsertBatch100-10             1        4528676083 ns/op         1834016 B/op      51549 allocs/op
//BenchmarkInsertBatch500-10             1        1111784166 ns/op         1971376 B/op      50173 allocs/op
//BenchmarkInsertBatch1000-10        86997             16423 ns/op              24 B/op          0 allocs/op
//BenchmarkInsertBatch5000-10     262566420                4.143 ns/op           0 B/op          0 allocs/op

const (
	dsn          = "username:password@tcp(localhost:3306)/dbname"
	totalRecords = 10000
)

//func BenchmarkInsertBatch1(b *testing.B)    { benchmarkInsertBatch(b, 1) }
func BenchmarkInsertBatch100(b *testing.B)  { benchmarkInsertBatch(b, 100) }
func BenchmarkInsertBatch500(b *testing.B)  { benchmarkInsertBatch(b, 500) }
func BenchmarkInsertBatch1000(b *testing.B) { benchmarkInsertBatch(b, 1000) }
func BenchmarkInsertBatch5000(b *testing.B) { benchmarkInsertBatch(b, 5000) }

// create table cloud.cloud_data_device_first_log
//(
//    id          int auto_increment
//        primary key,
//    device_id   varchar(100) not null comment '设备ID',
//    create_time bigint       not null comment '创建时间',
//    constraint device_id
//        unique (device_id)
//);

func benchmarkInsertBatch(b *testing.B, batchSize int) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 每次运行基准测试前清空表数据
	_, err = db.Exec("TRUNCATE TABLE cloud_data_device_first_log")
	if err != nil {
		log.Fatal(err)
	}

	b.ResetTimer()

	deviceIDCounter := 0

	for i := 0; i < totalRecords; i += batchSize {
		var records []string
		for j := 0; j < batchSize && i+j < totalRecords; j++ {
			deviceIDCounter++
			record := fmt.Sprintf("('%s', %d)", fmt.Sprintf("device%d", deviceIDCounter), time.Now().Unix())
			records = append(records, record)
		}

		// 批量插入
		insertBatch(db, records)
	}
}

func insertBatch(db *sql.DB, records []string) {
	var queryBuilder strings.Builder
	queryBuilder.WriteString("INSERT INTO cloud_data_device_first_log (device_id, create_time) VALUES ")
	queryBuilder.WriteString(strings.Join(records, ","))

	query := queryBuilder.String()
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
