package main

import (
	"github.com/go-martini/martini"
	"log"
	"net/http"
)

var installed = false

func main() {
	m := martini.Classic()

	m.Get("/install", func() string {
		installed = true
		return "i am the install or register page"
	})

	m.Get("/", func() string {
		return "i am the root"
	})

	m.Get("/index", func() string {
		return "i am the index page"
	})

	// 首先进入install 页面进行安装
	m.Use(MiddleWareTest)
	m.Run()
}

func MiddleWareTest(res http.ResponseWriter, req *http.Request) {
	log.Println(req.RequestURI)
	if installed == true || req.RequestURI == "/install" {
		log.Println("install page self")
	} else { //未安装、注册且访问路径不是 /install
		log.Println("redirected to install or register page")
		http.Redirect(res, req, "/install", 302)
		return
	}
}
