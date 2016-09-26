package main

import (
	"code.google.com/p/mahonia"
	"fmt"
)

func mai1n() {
	src := "编码转换内容内容"
	enc := mahonia.NewEncoder("GBK")
	output := enc.ConvertString(src)
	fmt.Println(output)
}

func main() {

	//"你好，世界！"的GBK编码

	testBytes := []byte{0xbf, 0xb5, 0xd0, 0xa1, 0xdc, 0xe7}

	var testStr string

	utfStr := "你好，世界！"

	var dec mahonia.Decoder

	var enc mahonia.Encoder

	testStr = string(testBytes)

	dec = mahonia.NewDecoder("gbk")

	if ret, ok := dec.ConvertStringOK(testStr); ok {

		fmt.Println("GBK to UTF-8: ", ret, " bytes:", testBytes)

	}

	fmt.Println(utfStr)

	enc = mahonia.NewEncoder("gbk")

	if ret, ok := enc.ConvertStringOK(utfStr); ok {

		fmt.Println("UTF-8 to GBK: ", ret, " bytes: ", []byte(ret))

	}

	return

}

func utf2gbk() {

	utfStr := "你好，世界！"

	var enc mahonia.Encoder

	enc = mahonia.NewEncoder("gbk")

	if ret, ok := enc.ConvertStringOK(utfStr); ok {

		fmt.Println("UTF-8 to GBK: ", ret, " bytes: ", []byte(ret))

	}

	return

}

func gbk2utf() {

	//"你好，世界！"的GBK编码

	testBytes := []byte{0xC4, 0xE3, 0xBA, 0xC3, 0xA3, 0xAC, 0xCA, 0xC0, 0xBD, 0xE7, 0xA3, 0xA1}

	var testStr string

	var dec mahonia.Decoder

	testStr = string(testBytes)

	dec = mahonia.NewDecoder("gbk")

	if ret, ok := dec.ConvertStringOK(testStr); ok {

		fmt.Println("GBK to UTF-8: ", ret, " bytes:", testBytes)

	}

	return

}
