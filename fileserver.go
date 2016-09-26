package main

import "net/http"

func main() {
	h := http.FileServer(http.Dir("d:\\"))
	http.ListenAndServe(":9999", h)
}
