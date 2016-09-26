package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"unicode"
)

func main() {
	content := readAll("DZDZ_143000000-DZDZ_142000000.sql")
	// fmt.Println(str)
	RemoveChineseChar(content)
}

var test = "（）！，。≡（： 中文"

func RemoveChineseChar(str string) {

	set := []*unicode.RangeTable{unicode.Han /*, unicode.P*/} // 中文汉字和 标点(含英文标点)
	// rs := []rune(str)

	for _, r := range str {
		// if !unicode.Is(unicode.Scripts["Han"], r) {
		// fmt.Print(string(r))
		if unicode.IsOneOf(set, r) { // 中文
			fmt.Print("H")
		} else if string(r) == "（" { // 中文标点 全角
			fmt.Print("H")
		} else if string(r) == "）" {
			fmt.Print("H")
		} else if string(r) == "！" {
			fmt.Print("H")
		} else if string(r) == "，" {
			fmt.Print("H")
		} else if string(r) == "。" {
			fmt.Print("H")
		} else if string(r) == "：" {
			fmt.Print("H")
		} else {
			fmt.Print(string(r))
		}

	}
}

func readAll(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}
