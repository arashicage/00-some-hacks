package main

import (
	"fmt"

	"gopkg.in/macaron.v1"
)

func main() {
	m := macaron.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	// m.Run(8080)
	m.Run() //default 4000
	fmt.Println("hello world")
}
