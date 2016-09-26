package main

import (
	"gopkg.in/macaron.v1"
)

func main() {
	m := macaron.Classic()
	m.Get("/", func() string {
		return "Hello from site 1!"
	})

	m.Get("/site1", func() string {
		return "Hello from site 1! with url /site1"
	})

	m.Run(8001)
}
