package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	neo4j "github.com/davemeehan/Neo4j-GO"
)

// HOST USER PASSWORD DBNAME
const (
	HOST          = "http://127.0.0.1:7474/db/data"
	USER          = "neo4j"
	PASSWORD      = "root"
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
	neo, err := neo4j.NewNeo4j(HOST, USER, PASSWORD)
	checkErr(err)
	// iNode1 := map[string]string{
	// 	"uuid":   "sushdausdhka",
	// 	"name":   "你",
	// 	"pinyin": "ni",
	// }
	// iNode2 := map[string]string{
	// 	"uuid":   "susadshdausdhka",
	// 	"name":   "好",
	// 	"pinyin": "hao",
	// }

	// data1, _ := neo.CreateNode(iNode1)
	// data2, _ := neo.CreateNode(iNode2)
	// log.Printf("\nNode data: %v\n", data1)
	// log.Printf("\nNode data: %v\n", data2)
	// real := map[string]string{
	// 	"isB": "false",
	// 	"isT": "true",
	// }
	// neo.CreateRelationship(3, 4, real, "BT")
	// neo.Traverse(10, "node", "3", "", nil, 5, nil, nil)
	data := map[string]string{
		"s": "b",
	}
	dd, _ := neo.CreateNode(data)
	log.Printf("\nNode data: %v\n", dd)

	// checkErr(err)

	// getCiXing("你好")
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
