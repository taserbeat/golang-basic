package main

import (
	"fmt"
	"main/test"
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
}
