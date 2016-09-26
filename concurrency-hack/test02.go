package main

import "fmt"

type Sushi string

func main() {

	// 译注：要想运行该示例，需要先定义类型 Sushi，如type Sushi string
	var ch <-chan Sushi = Producer()
	for s := range ch {
		fmt.Println("Consumed", s)
	}
}

func Producer() <-chan Sushi {
	ch := make(chan Sushi)
	go func() {
		ch <- Sushi("海老握り")  // Ebi nigiri
		ch <- Sushi("鮪とろ握り") // Toro nigiri
		close(ch)
	}()
	return ch
}
