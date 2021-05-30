package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"main/test"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"time"
)

// json
type A struct {
}

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	A         A         `json:"A"`
}

func (user User) CustomMarshalJson() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Id        int
		Name      string
		Email     string
		CreatedAt time.Time
		A         A
	}{
		Id:        user.Id,
		Name:      "Mr " + user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		A:         user.A,
	})

	return v, err
}

/* sortパッケージ */
type Entry struct {
	Key   string
	Value int
}

/* net/http (server) */
type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
}

func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("public/top.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, "this is top")
}

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

	/* json */

	user1 := User{
		Id:        1,
		Name:      "John",
		Email:     "john@example.com",
		CreatedAt: time.Now(),
	}

	// marshal: 構造体からJSONに変換
	jsonBs, err := json.Marshal(user1)
	if err != nil {
		log.Fatal(err)
	}
	// {"id":1,"name":"John","email":"john@example.com","createdAt":"2021-05-30T17:10:21.824778+09:00","A":{}}
	fmt.Println(string(jsonBs))

	// unarshal: JSONから構造体に変換
	user2 := new(User)
	if err := json.Unmarshal(jsonBs, &user2); err != nil {
		fmt.Println(err)
	}

	fmt.Println(user2) // &{1 John john@example.com 2021-05-30 17:19:40.08609 +0900 JST {}}

	// Marshalのカスタマイズ
	customJsonBs, err := user1.CustomMarshalJson()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(customJsonBs)) // {"Id":1,"Name":"Mr John","Email":"john@example.com","CreatedAt":"2021-05-30T17:37:09.051518+09:00","A":{}}

	/* sortパッケージ */
	numbers := []int{5, 3, 7, 1, 6, 9, 2, 8, 4}
	strings := []string{"j", "z", "a"}

	fmt.Println(numbers, strings) // [5 3 7 1 6 9 2 8 4] [a z j]

	sort.Ints(numbers)
	sort.Strings(strings)

	fmt.Println(numbers, strings) // [1 2 3 4 5 6 7 8 9] [a j z]

	entries1 := []Entry{
		{"A", 20},
		{"J", 10},
		{"n", 50},
		{"I", 80},
		{"t", 60},
		{"B", 60},
		{"E", 40},
		{"p", 40},
	}

	entries2 := make([]Entry, len(entries1))
	copy(entries2, entries1)

	// Entry型のスライスをKeyでソート
	fmt.Println("Entry型のスライスをKeyで降順ソート")
	fmt.Println(entries1) // [{A 20} {J 10} {n 50} {I 80} {t 60} {B 60} {E 40} {p 40}]

	sort.Slice(entries1, func(i, j int) bool { return entries1[i].Key < entries1[j].Key })

	fmt.Println(entries1) // [{A 20} {B 60} {E 40} {I 80} {J 10} {n 50} {p 40} {t 60}]]

	// Entry型のスライスをValueでStableSort (安定ソート)
	fmt.Println("StableSort (安定ソート)")
	fmt.Println(entries2) // [{A 20} {J 10} {n 50} {I 80} {t 60} {B 60} {E 40} {p 40}]

	sort.SliceStable(entries2, func(i, j int) bool { return entries2[i].Value < entries2[j].Value })

	fmt.Println(entries2) // [{J 10} {A 20} {E 40} {p 40} {n 50} {t 60} {B 60} {I 80}]

	/* contextパッケージ */
	// ある処理にかかる時間が長い場合に途中で処理を中断することを可能にする

	/* net/urlパッケージ */
	// URL文字列を処理するパッケージである。

	// URLを解析
	parser, _ := url.Parse("http://example.com/search?a=1&b=2#top")
	fmt.Println(parser.Scheme)   // http
	fmt.Println(parser.Host)     // example.com
	fmt.Println(parser.Path)     // /search
	fmt.Println(parser.RawQuery) // a=1&&b=2
	fmt.Println(parser.Fragment) // top

	fmt.Println(parser.Query()) // map[a:[1] b:[2]]

	// URLを生成
	encoder := &url.URL{}
	encoder.Scheme = "https"
	encoder.Host = "google.com"
	q := encoder.Query()
	q.Set("q", "Golang")

	encoder.RawQuery = q.Encode()

	fmt.Println(encoder) // https://google.com?q=Golang

	/* net/httpパッケージ (client) */
	// HTTPクライアントのためのパッケージ

	// GETメソッド
	getResponse, _ := http.Get("https://example.com")

	fmt.Println(getResponse.StatusCode) // 200
	fmt.Println(getResponse.Proto)      // HTTP/2.0

	fmt.Println(getResponse.Header["Date"])         // [Sun, 30 May 2021 11:43:51 GMT]
	fmt.Println(getResponse.Header["Content-Type"]) // [text/html; charset=UTF-8]

	fmt.Println(getResponse.Request.Method) // GET
	fmt.Println(getResponse.Request.URL)    // https://example.com

	body, _ := ioutil.ReadAll(getResponse.Body)
	fmt.Println(string(body)) // HTTPレスポンスのbodyを文字列として出力

	// POSTメソッド
	values := url.Values{}

	values.Add("id", "1")
	values.Add("message", "メッセージ")
	fmt.Println(values.Encode())

	postResponse, err := http.PostForm("https://example.com/", values)
	if err != nil {
		log.Fatal(err)
	}

	postResponseBody, _ := ioutil.ReadAll(postResponse.Body)
	fmt.Println(string(postResponseBody))

	/* net/http (server) */

	http.HandleFunc("/top", top)      // curl http://localhost:8080/top   -> this is top
	http.ListenAndServe(":8080", nil) // curl http://localhost:8080   -> 404

}
