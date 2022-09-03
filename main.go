package main

import (
	"ipumpkin/handler"

	"github.com/flamego/flamego"
)

func main() {
	f := flamego.New()

	// 渲染中间件
	f.Use(flamego.Renderer())
	f.Group("/",
		func() {
			f.Get("test", handler.Fyhtest)
		})
	f.Run("127.0.0.1", "2630")

}
