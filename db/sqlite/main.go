package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// database + sqlite3

const SQL_PATH = "./sql/example.sql"

var Db *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	// sqliteのファイルが存在していたら削除することでテーブルをリセット
	_, err := os.Stat(SQL_PATH)
	isExist := !os.IsNotExist(err)
	if isExist {
		err = os.Remove(SQL_PATH)
		if err != nil {
			log.Fatalln(err)
		}
	}

	// テーブル作成
	Db, _ := sql.Open("sqlite3", SQL_PATH)
	defer Db.Close()

	cmd := `CREATE TABLE IF NOT EXISTS persons(
    name STRING,
    age INT
  )`

	_, err = Db.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	// データの追加
	cmd = "INSERT INTO persons (name, age) VALUES (?, ?), (?, ?), (?, ?)"
	_, err = Db.Exec(cmd, "taro", 20, "hanako", 15, "ichiro", 40)
	if err != nil {
		log.Fatalln(err)
	}

	// データの更新
	cmd = "UPDATE persons SET age = ? WHERE name = ?"
	_, err = Db.Exec(cmd, 30, "taro")
	if err != nil {
		log.Fatalln(err)
	}

	// データの取得
	cmd = "SELECT * FROM persons WHERE age = ?"
	// QueryRaw 1レコード取得
	record := Db.QueryRow(cmd, 30)
	var person Person
	err = record.Scan(&person.Name, &person.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Not Found record")
		} else {
			log.Println(err)
		}
	}
	fmt.Println("---1レコード取得---")
	fmt.Println(person.Name, person.Age)

	// 複数レコードの取得
	cmd = `SELECT * FROM persons`
	rows, _ := Db.Query(cmd) // Queryで条件に合うレコードを全て取得
	defer rows.Close()

	var persons []Person
	for rows.Next() {
		var p Person
		err = rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		persons = append(persons, p)
	}

	fmt.Println("---複数レコード取得---")
	for _, p := range persons {
		fmt.Println(p.Name, p.Age)
	}

	// データの削除
	cmd = `DELETE FROM persons WHERE name = ?`
	_, err = Db.Exec(cmd, "ichiro")
	if err != nil {
		log.Fatalln(err)
	}

}
