package main

import (
	"ipumpkin/config"
	"ipumpkin/handler"

	"github.com/flamego/flamego"
)

func main() {
	mongodb := handler.InitDB(config.Config())
	f := flamego.New()
	// 渲染中间件
	f.Use(flamego.Renderer())
	f.Map(mongodb)
	f.Group("/",
		func() {
			f.Get("test", handler.Fyhtest)
			f.Get("testmongo", handler.FyhMongo)
		})
	f.Run("127.0.0.1", "2630")

}
