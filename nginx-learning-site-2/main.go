package main

import (
	"gopkg.in/macaron.v1"
)

func main() {
	m := macaron.Classic()
	m.Get("/", func() string {
		return "Hello from site 2!"
	})

	m.Get("/site2", func() string {
		return "Hello from site 2! with url /site2"
	})

	m.Run(8002)
}
