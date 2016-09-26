package main

import (
	"fmt"
	"strconv"
)

func main() {
	err := defer_test()

	fmt.Println("main ", err)
}

func defer_test() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	age, err := strconv.Atoi("five")
	if err != nil {
		fmt.Println("发生错误", age)
		panic(err)
	}

	return err
}
