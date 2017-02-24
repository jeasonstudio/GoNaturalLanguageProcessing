package main

import (
	"database/sql"
	"fmt"
)

// HOST USER PASSWORD DBNAME
const (
	HOST     = "127.0.0.1"
	PORT     = "3306"
	USER     = "root"
	PASSWORD = ""
	DBNAME   = "naturl_language_process"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", USER+":@tcp("+HOST+":"+PORT+")/"+DBNAME+"?charset=utf8")
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.Ping()
}

func main() {
	fmt.Println(HOST)
	rows, _ := db.Query("SELECT * FROM main_single_word WHERE word_id = '4'")
	defer rows.Close()

	for rows.Next() {
		var word_id, word_name, word_pinyin string
		rows.Scan(&word_id, &word_name, &word_pinyin)
		fmt.Println("word_id:", word_id, "word_name:", word_name, "word_pinyin:", word_pinyin)
	}
}
