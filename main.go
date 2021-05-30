package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"main/test"
	"math"
	"math/rand"
	"os"
	"regexp"
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

	/* regexpパッケージ */
	// Goの正規表現の基本
	match, _ := regexp.MatchString("A", "ABC")
	fmt.Println(match) // true

	// Compile
	re1, _ := regexp.Compile("A")
	match = re1.MatchString("ABC")
	fmt.Println(match) // true

	// MustCompile
	re2 := regexp.MustCompile("A")
	match = re2.MatchString("BCD")
	fmt.Println(match) // false

	// 正規表現のフラグ
	/*
		フラグオプション一覧

		i 大文字小文字を区別しない
		m マルチラインモード(^と&が文頭、文末に加えて行頭、行末にマッチ)
		s .が\nにマッチ
		U 最小マッチへの変換 (xxはx+?へ、x+はx+?へ)
	*/

	re3 := regexp.MustCompile(`(?i)abc`) // (?)にフラグオプションをつけることで正規表現のルールを適用できる
	match = re3.MatchString("ABC")
	fmt.Println(match) // true

	// 幅を持たない正規表現のパターン
	/*
		パターン一覧

		^ 文頭 (mフラグが有効な場合は行頭にも)
		$ 文末 (mフラグが有効な場合は行頭にも)
		\A 文頭
		\z 文末
		\b ASCIIに夜ワード協会
		\B 非ASCIIによるワード協会

	*/

	re4 := regexp.MustCompile(`^ABC$`)
	match = re4.MatchString("ABC")
	fmt.Println(match) // true

	match = re4.MatchString("   ABC   ")
	fmt.Println(match) // false

	// 繰り返しを表す正規表現
	/*
		繰り返しのパターン

		x* 0回以上繰り返すx (最大マッチ)
		x+ 1回以上繰り返すx (最大マッチ)
		x? 0回以上1回以下繰り返すx
		x{n,m} n回以上m回以下繰り返すx (最大マッチ)
		x{n, } n回以上繰り返すx (最大マッチ)
		x{n} n回繰り返すx (最大マッチ)

		x*? 0回以上繰り返すx (最小マッチ)
		x+? 1回以上繰り返すx (最小マッチ)
		x?? 0回以上1回以下繰り返すx (0回優先)
		x{n,m}? n回以上m回以下繰り返す (最小マッチ)
		x{n, }? n回以上繰り返すx (最小マッチ)
		x{n}? n回繰り返すx (最小マッチ)
	*/

	re5 := regexp.MustCompile("a+b*")
	fmt.Println(re5.MatchString("a"))  // true
	fmt.Println(re5.MatchString("b"))  // false
	fmt.Println(re5.MatchString("ab")) // true
	fmt.Println(re5.MatchString("aa")) // true
	fmt.Println(re5.MatchString("bb")) // false

	// 正規表現の文字クラス
	re6 := regexp.MustCompile(`[XYZ]`)
	fmt.Println(re6.MatchString("Y")) // true
	fmt.Println(re6.MatchString("V")) // false

	re7 := regexp.MustCompile(`^[0-9A-Za-z_]{3}$`) // 英数字と_に3回マッチする
	fmt.Println(re7.MatchString("ABC"))            // true
	fmt.Println(re7.MatchString("abcdefg"))        // false

	re8 := regexp.MustCompile(`[^0-9A-Za-z_]`) // 英数字と_以外にマッチする
	fmt.Println(re8.MatchString("ABC"))        // false
	fmt.Println(re8.MatchString("あ"))          // true

	// 正規表現にマッチした文字列の取得
	re9 := regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`)

	fs1 := re9.FindString("AAAABCXYZZZZ")
	fmt.Println(fs1) // ABCXYZ

	fs2 := re9.FindAllString("ABCXYZABCXYZ", -1) // 第2引数で取得する文字列スライスの数を指定する。-1を指定するとマッチした全てを取得する。
	fmt.Println(fs2)                             // [ABCXYZ ABCXYZ]

	/* cryptパッケージ */
	// MD5ハッシュ値を生成
	// 任意の文字列からMD5のハッシュ値を生成する処理例
	hashGen := md5.New()
	io.WriteString(hashGen, "ABCDE")

	// ハッシュ値のバイト配列を取得する
	fmt.Println(hashGen.Sum(nil)) // [46 205 222 57 89 5 29 145 63 97 177 69 121 234 19 109]

	hashHex := fmt.Sprintf("%x", hashGen.Sum(nil)) // ハッシュ値のバイト配列 -> 16進数の文字列を得る
	fmt.Println(hashHex)                           // 2ecdde3959051d913f61b14579ea136d

}
