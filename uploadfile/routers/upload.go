package routers

import (
	"fmt"
	"gopkg.in/macaron.v1"
	"io"
	"os"
)

func ForgotPasswd2(ctx *macaron.Context) {
	r := ctx.Req
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("fileToUpload")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	ctx.Write([]byte("hello"))
}
