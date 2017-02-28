package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"gopkg.in/jmcvetta/neoism.v1"
)

// HOST USER PASSWORD DBNAME
const (
	HOST          = "http://neo4j:root@127.0.0.1:7474/db/data"
	USER          = "neo4j"
	PASSWORD      = "root"
	DBNAME        = "naturl_language_process"
	LANGUAGECLOUD = "f9Z9q9U6uUbDNfmEkEjHKAFbMhBBcftx0N5y7W6f"
	URL           = "http://api.ltp-cloud.com/analysis/"
)

func main() {
	fmt.Println(getCiXing("我"))
	// dat, err := ioutil.ReadFile("rescourse/single_word.tsv")
	// checkErr(err)
	// res := strings.Split(string(dat), "\n")
	// for i := 0; i < 5; i++ {
	// 	name := strings.Split(res[i], "\t")[1]
	// 	poSpeech := getCiXing(name)
	// 	pinyin := strings.Split(res[i], "\t")[2]
	// 	label := strings.ToUpper(strings.Split(pinyin, "")[0])
	// 	fmt.Println(name, poSpeech, pinyin, (label))
	// 	time.Sleep(1 * time.Second)
	// }
}

func connectToNeo4j() {
	myNeo, err := neoism.Connect(HOST)
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

	fNode, _ := myNeo.CreateNode(neoism.Props{"name": "Captain Kirk"})
	defer fNode.Delete() // Deferred clean up
	fNode.AddLabel("Person")
	log.Printf("\nNode data: %v\n", fNode)

	res0 := []struct {
		N neoism.Node // Column "n" gets automagically unmarshalled into field N
	}{}
	cq0 := neoism.CypherQuery{
		Statement: "CREATE (n:Person {name: {name}}) RETURN n",
		// Use parameters instead of constructing a query string
		Parameters: neoism.Props{"name": "Jeason"},
		Result:     &res0,
	}
	myNeo.Cypher(&cq0)

	// checkErr(err)
	// getCiXing("你好")
}

func getCiXing(text string) string {
	resp, err := http.Get(URL + "?api_key=" + LANGUAGECLOUD + "&text=" + text + "&pattern=pos&format=plain")
	checkErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	// fmt.Println(string(body))
	resArr := strings.Split(string(body), "_")
	if len(resArr) > 1 {
		return resArr[1]
	} else {
		return " "
	}
	// for i := range resArr {
	// 	thisName := strings.Split(resArr[i], "_")
	// 	resArr[i] = thisName[1]
	// }
	// whatYouWant := strings.Join(resArr, "/")
	// fmt.Println("Normal Text is", string(body), "CiXing is", whatYouWant, "Success!!")

	// return resArr[1]
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

// func getcx(text string) {
// 	connection := flag.Int("c", 10, "-c N")
// 	timeout := flag.Int("o", 5, "-o N")
// 	timeover := flag.Int("t", 5, "-t N")
// 	printresult := flag.Bool("p", false, "-p false")
// 	method := flag.String("m", "GET", "-m GET")
// 	url := URL + "?api_key=" + LANGUAGECLOUD + "&text=" + text + "&pattern=pos&format=plain"
// 	flag.Parse()
// 	var Count int32
// 	defer func() {
// 		if !*printresult {
// 			fmt.Println("成功响应：", Count)
// 		}
// 	}()
// 	T := time.Tick(time.Duration(*timeover) * time.Second)
// 	var result chan string = make(chan string, 10)
// 	t := time.Duration(*timeout) * time.Second
// 	Client := http.Client{Timeout: t}
// 	for i := 0; i < *connection; i++ {
// 		go func() {
// 			req, _ := http.NewRequest(*method, url, nil)
// 			resp, _ := Client.Do(req)
// 			defer resp.Body.Close()
// 			if resp.StatusCode == 200 {
// 				b, _ := ioutil.ReadAll(resp.Body)
// 				// fmt.Println(string(b))
// 				result <- string(b)
// 				atomic.AddInt32(&Count, int32(1))
// 			}
// 		}()
// 	}
// 	for {
// 		select {
// 		case x := <-result:
// 			if *printresult {
// 				fmt.Print(x)
// 			}
// 		case <-T:
// 			return
// 		}
// 	}
// }
