package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
)

// HOST USER PASSWORD DBNAME
const (
	HOST          = "127.0.0.1"
	PORT          = "3306"
	USER          = "root"
	PASSWORD      = ""
	DBNAME        = "naturl_language_process"
	LANGUAGECLOUD = "f9Z9q9U6uUbDNfmEkEjHKAFbMhBBcftx0N5y7W6f"
	URL           = "http://api.ltp-cloud.com/analysis/"
)

// var db *sql.DB

// func init() {
// 	db, _ = sql.Open("mysql", USER+"@tcp("+HOST+":"+PORT+")/"+DBNAME+"?charset=utf8")
// 	db.SetMaxOpenConns(2000)
// 	db.SetMaxIdleConns(1000)
// 	db.Ping()
// }

func main() {
	db, _ := gorm.Open("mysql", USER+"@tcp("+HOST+":"+PORT+")/"+DBNAME+"?charset=utf8")
	defer db.Close()

	var user string

	x := db.Where("word_id = ?", "4").Find(&user)

	fmt.Println(x, user)

	getCiXing("你好")
}

func getCiXing(text string) string {
	resp, err := http.Get(URL + "?api_key=" + LANGUAGECLOUD + "&text=" + text + "&pattern=pos&format=plain")
	checkErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	resArr := strings.Split(string(body), " ")
	for i := range resArr {
		thisName := strings.Split(resArr[i], "_")
		resArr[i] = thisName[1]
	}
	whatYouWant := strings.Join(resArr, "/")
	fmt.Println("Normal Text is", string(body), "CiXing is", whatYouWant, "Success!!")
	return whatYouWant
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
