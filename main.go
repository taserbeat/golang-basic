package main

import (
	"fmt"
	"main/test"
	"math"
	"math/rand"
	"time"
)

// テスト

func main() {
	fmt.Println(test.IsOne(1))
	fmt.Println(test.IsOne(0))

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(test.Average(s))

	/* timeパッケージ */
	// 時間の生成
	nowTime := time.Now()
	fmt.Println(nowTime)

	// 指定した時間を生成
	specificTime := time.Date(2020, 6, 12, 13, 14, 15, 0, time.Local)
	fmt.Println(specificTime)

	/* mathパッケージ */
	// 最大値と最小値
	maxStr := fmt.Sprintf("%dと%dの最大値: %d", 9, 10, int(math.Max(9, 10)))
	fmt.Println(maxStr)
	minStr := fmt.Sprintf("%dと%dの最小値: %d", 1, 2, int(math.Min(1, 2)))
	fmt.Println(minStr)

	// 小数点以下を切り捨てる
	fmt.Println(math.Trunc(3.14))  // 3
	fmt.Println(math.Trunc(-3.14)) // -3

	/* randパッケージ */
	// 乱数生成器にシードを設定
	rand.Seed(256)

	// 乱数を生成
	fmt.Println(rand.Float64()) // 0.813527291469711
	fmt.Println(rand.Float64()) // 0.5598026045235738
	fmt.Println(rand.Float64()) // 0.6695717783859498

	// 現在の時刻をシードに使った疑似乱数の生成
	rand.Seed(time.Now().UnixNano())

	fmt.Println(rand.Int())     // int型の疑似乱数
	fmt.Println(rand.Intn(100)) // 0 ~ 99のint型疑似乱数
	fmt.Println(rand.Int63())   // int64型の疑似乱数

}
