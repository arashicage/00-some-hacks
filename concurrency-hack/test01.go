package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		ch <- "Hello!"
		close(ch)
	}()
	fmt.Println("1", <-ch) // 输出字符串"Hello!"
	fmt.Println("2", <-ch) // 输出零值 - 空字符串""，不会阻塞
	fmt.Println("3", <-ch) // 再次打印输出空字符串"" v, ok := <-ch // 变量v的值为空字符串""，变量ok的值为false
}
