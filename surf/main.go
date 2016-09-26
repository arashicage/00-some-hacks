package main

import (
	"fmt"
	"github.com/headzoo/surf"
)

func main() {
	bow := surf.NewBrowser()
	err := bow.Open("https://github.com/arashicage/surf")
	if err != nil {
		panic(err)
	}

	// Outputs: "The Go Programming Language"
	fmt.Println(bow.Title())
}
