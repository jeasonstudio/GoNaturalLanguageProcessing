package controller

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	pinyin "github.com/mozillazg/go-pinyin"
)

// ControMain nothing
func ControMain() {
	theStr, _ := readAll("rescourse/CIHUI2.txt")
	arrSlice := strings.Split(string(theStr), "\n")
	// fmt.Println(arrSlice)
	for i := range arrSlice {
		a := pinyin.NewArgs()
		a.Style = 8
		tt := strings.Split(arrSlice[i], "	")[0]
		k := pinyin.Pinyin(tt, a)
		var m []string
		for i := 0; i < len(k); i++ {
			m = append(m, k[i][0])
		}
		b := strings.Join(m, "/")
		fmt.Println("INSERT INTO naturl_language_process.main_single_terms (terms_name, terms_pinyin) VALUES ('" + tt + "', '" + b + "');")
		// fmt.Println(arrSlice[i], b)
	}
}

func readAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}
