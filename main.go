package main

import (
	"ipumpkin/common"
	"ipumpkin/config"
	"ipumpkin/handlers"

	"github.com/flamego/flamego"
)

func main() {
	mongodb := handlers.InitDB(config.Config())
	cfgProvider := config.LoadConfigProvider()
	logger := common.NewLogger(cfgProvider)
	f := flamego.New()
	// 渲染中间件
	go handlers.DockerOperate(cfgProvider, mongodb, logger)
	f.Use(flamego.Renderer())
	f.Map(mongodb)
	f.Map(logger)
	f.Map(cfgProvider)
	f.Group("/",
		func() {
			f.Get("mongo", handlers.FyhMongo)
		})
	f.Run(cfgProvider.GetString("server.host"), cfgProvider.GetString("server.port"))

}
