package main

import (
	"gopkg.in/macaron.v1"
	"uploadfile/routers"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	m := macaron.Classic()

	m.Get("/", func() string {
		return "Hello world!"
	})

	m.Post("/upload", routers.ForgotPasswd2)

	m.Run()
}
