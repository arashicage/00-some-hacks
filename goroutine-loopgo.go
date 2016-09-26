package main

import (
	"fmt"
	"time"
)

func main() {
	sem := make(chan int, 10)
	//FOR循环体
	for i := 0; i < 10; i++ {
		//建立协程
		go func(i int) {
			fmt.Println(i, time.Now())
			//计数
			sem <- 0
		}(i)
	}
	// 等待循环结束
	for i := 0; i < 10; i++ {
		<-sem
	}
}
