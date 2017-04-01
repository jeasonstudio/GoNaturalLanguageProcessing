package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

func init() {
}

func main() {
	myNeo, err := neoism.Connect(HOST)
	checkErr(err)
	CQL := `
		
	`
	addCQL(myNeo, CQL)
}

func luru() {
	myNeo, err := neoism.Connect(HOST)
	checkErr(err)
	// CQL := `
	// 	START a=node(*),b=node(*)
	// 	WHERE a.name='你' AND b.name='好'
	// 	CREATE (a)-[n:connectTo{termsNum:'0',relPinyin:'hhhh',thisTermNum:'2'}]->(b)
	// 	RETURN n
	// `

	// addCQL(myNeo, CQL)
	dat, err := ioutil.ReadFile("rescourse/terms.tsv")
	checkErr(err)

	res := strings.Split(string(dat), "\n")

	for k := 0; k < len(res); k++ {
		// fmt.Println(strings.Split(res[k], "\t")[1])
		term := strings.Split(strings.Split(res[k], ",")[1], "")
		relPinyin := strings.Split(res[k], ",")[2]
		if len(term) < 1 {
			fmt.Println("No Create Relationship")
		} else {
			for j := 0; j < len(term)-1; j++ {
				fN := term[j]
				sN := term[j+1]
				termsLen := strconv.Itoa(len(term))
				fromTo := strconv.Itoa(j) + `-` + strconv.Itoa(j+1)
				CQL := "START a=node(*),b=node(*) WHERE a.name='" + fN + "' AND b.name='" + sN + "' CREATE (a)-[n:connectTo{termsLen:'" + termsLen + "',relPinyin:'" + relPinyin + "',fromTo:'" + fromTo + "'}]->(b) RETURN n"
				addCQL(myNeo, CQL)
				fmt.Println(CQL)
			}
		}

		// fmt.Println(term, relPinyin)
	}

}

func addCQL(r *neoism.Database, c string) {

	cq1 := neoism.CypherQuery{
		Statement: c,
	}
	err := r.Cypher(&cq1)
	checkErr(err)
}

func createNodes() {
	myNeo, err := neoism.Connect(HOST)
	checkErr(err)
	// fmt.Println(getCiXing("我"))

	dat, err := ioutil.ReadFile("rescourse/single_word.tsv")
	checkErr(err)
	res := strings.Split(string(dat), "\n")
	for i := 0; i < len(res); i++ {
		name := strings.Split(res[i], "\t")[1]
		// poSpeech := getCiXing(name)
		pinyin := strings.Split(res[i], "\t")[2]
		label := strings.ToUpper(strings.Split(pinyin, "")[0])

		fmt.Println(name + "," + pinyin + "," + (label))

		res0 := []struct {
			N neoism.Node // Column "n" gets automagically unmarshalled into field N
		}{}
		cq0 := neoism.CypherQuery{
			Statement:  "CREATE (n:" + label + " {name: {name},pinyin: {pinyin}}) RETURN n",
			Parameters: neoism.Props{"name": name, "pinyin": pinyin},
			Result:     &res0,
		}
		myNeo.Cypher(&cq0)
		n1 := res0[0].N // Only one row of data returned
		n1.Db = myNeo   // Must manually set Db with objects returned from Cypher query
	}
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
