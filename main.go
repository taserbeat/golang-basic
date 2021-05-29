package main

import (
	"fmt"
	"log"
	"main/test"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
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

	/* logパッケージ */
	// ログの出力先を標準出力に変更
	log.SetOutput(os.Stdout)

	log.Println("this is log.")

	// loggerの作成
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("message")

	/* strconvパッケージ */
	// 整数 -> 文字列にの変換
	strFromInt1 := strconv.FormatInt(-100, 10)
	fmt.Printf("%v, %T\n", strFromInt1, strFromInt1)

	// もっと簡単にint -> stringの変換ができる
	strFromInt2 := strconv.Itoa(100)
	fmt.Printf("%v, %T\n", strFromInt2, strFromInt2)

	// 文字列 ->boolの変換
	boolFromStr1, _ := strconv.ParseBool("true")
	fmt.Printf("%v, %T\n", boolFromStr1, boolFromStr1)

	boolFromStr2, ok := strconv.ParseBool("false")
	if ok != nil {
		fmt.Println("Convert Error")
	}
	fmt.Printf("%v, %T\n", boolFromStr2, boolFromStr2)

	// 文字列 -> intの変換
	intFromStr1, _ := strconv.ParseInt("12345", 10, 0) // 第3引数でint型の精度を指定する。0の場合はGoのint型の精度が設定される。
	fmt.Printf("%v, %T\n", intFromStr1, intFromStr1)

	// もっと簡単に文字列 -> intの変換ができる
	intFromStr2, _ := strconv.Atoi("123")
	fmt.Printf("%v, %T\n", intFromStr2, intFromStr2)

	// 文字列 -> floatの変換
	floatFromStr1, _ := strconv.ParseFloat("3.14", 64)
	fmt.Printf("%v, %T\n", floatFromStr1, floatFromStr1)

	/* stringsパッケージ */
	// 文字列を結合する
	fruits := strings.Join([]string{"banana", "watermelon", "apple"}, ", ")
	fmt.Println(fruits) // "banana, watermelon, apple"

	// 文字列に含まれる部分文字列を検索する
	index1 := strings.Index("ABCDEF", "C")
	index2 := strings.Index("ABCABC", "C")
	fmt.Println(index1, index2) // 2 2

}
