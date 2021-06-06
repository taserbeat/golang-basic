package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB
var err error

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, err = sql.Open("postgres", "host=127.0.0.1 port=15432 user=root password=postgres dbname=test_db sslmode=disable")

	defer Db.Close()

	if err != nil {
		log.Panicln(err)
	}

	// C
	cmd := "INSERT INTO persons (name, age) VALUES ($1, $2)"
	nancy := Person{Name: "Nancy", Age: 20}

	_, err = Db.Exec(cmd, nancy.Name, nancy.Age)
	if err != nil {
		log.Fatalln(err)
	}

	// R
	// QueryRow でレコードを1件取得
	cmd = "SELECT * FROM persons where age = $1"
	row := Db.QueryRow(cmd, 20)
	var p Person
	err = row.Scan(&p.Name, &p.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			// データが存在しない場合
			log.Println("Data Not Found")
		} else {
			// それ以外のエラー
			log.Println(err)
		}
	}
	fmt.Println("// レコードを1件取得")
	fmt.Println(p.Name, p.Age)

	// Query で条件に合うもの全てを取得
	cmd = "SELECT * FROM persons"
	rows, _ := Db.Query(cmd)
	defer rows.Close()
	var pp []Person
	// 取得したデータをループでスライスに追加 for rows.Next()
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		// 1つづつエラーハンドリング
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}

	// まとめてエラーハンドリングバージョン
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}
	// 表示
	fmt.Println("// 複数取得")
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}

	// U
	cmd = "UPDATE persons SET age = $1 WHERE name = $2"
	_, err := Db.Exec(cmd, 25, nancy.Name)
	if err != nil {
		log.Fatalln(err)
	}

	// D
	cmd = "DELETE FROM persons WHERE name = $1"
	_, err = Db.Exec(cmd, nancy.Name)
	if err != nil {
		log.Fatalln(err)
	}

}
