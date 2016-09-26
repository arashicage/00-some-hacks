package main

import (
	"github.com/go-macaron/cache"
	"github.com/go-macaron/captcha"
	"gopkg.in/macaron.v1"

	"fmt"
)

func main() {
	m := macaron.Classic()
	m.Use(cache.Cacher())
	m.Use(captcha.Captchaer())

	m.Use(macaron.Renderer(macaron.RenderOptions{
		Directory:  "templates",
		Extensions: []string{".tmpl", ".html"},
	}))

	m.Use(captcha.Captchaer(captcha.Options{
		// 获取验证码图片的 URL 前缀，默认为 "/captcha/"
		URLPrefix: "/captcha/",
		// 表单隐藏元素的 ID 名称，默认为 "captcha_id"
		FieldIdName: "captcha_id",
		// 用户输入验证码值的元素 ID，默认为 "captcha"
		FieldCaptchaName: "captcha",
		// 验证字符的个数，默认为 6
		ChallengeNums: 4,
		// 验证码图片的宽度，默认为 240 像素
		Width: 120,
		// 验证码图片的高度，默认为 80 像素
		Height: 40,
		// 验证码过期时间，默认为 600 秒
		Expiration: 600,
		// 用于存储验证码正确值的 Cache 键名，默认为 "captcha_"
		CachePrefix: "captcha_",
	}))

	m.Get("/", func(ctx *macaron.Context) {
		ctx.HTML(200, "hello")
	})

	m.Post("/x", func(ctx *macaron.Context, cpt *captcha.Captcha) string {
		if cpt.VerifyReq(ctx.Req) {
			fmt.Println("111")
			return "valid captcha"
		}
		fmt.Println("222")
		return "invalid captcha"
	})

	m.Run(8080)
}
