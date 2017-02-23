package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	pinyin "github.com/mozillazg/go-pinyin"
)

// HOST USER PASSWORD DBNAME
const (
	HOST     = "127.0.0.1"
	USER     = "root"
	PASSWORD = ""
	DBNAME   = "naturl_language_process"
)

func main() {
	theStr, _ := readAll("rescourse/CIHUI1.txt")
	arrSlice := strings.Split(string(theStr), "	")
	// fmt.Println(arrSlice)
	for i := range arrSlice {
		a := pinyin.NewArgs()
		a.Style = 8
		k := pinyin.Pinyin(arrSlice[i], a)
		var m []string
		for i := 0; i < len(k); i++ {
			m = append(m, k[i][0])
		}
		b := strings.Join(m, "/")
		// fmt.Println(i, arrSlice[i], strings.Join(b, "/"))
		fmt.Println(arrSlice[i], b)
	}
}

func readAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}
