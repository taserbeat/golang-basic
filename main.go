package main

import (
	"fmt"
	"main/test"
	"math"
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

}
