package server

import (
	"bms/src/server/route"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/spf13/viper"
)

func SetupServer() {
	isDevelopment := viper.GetBool("isDevelopment")

	app := iris.New()

	if isDevelopment {
		//iris自带日志文件
		app.Logger().SetLevel("debug")
		//设置recover从panics恢复，设置log记录
		//app.Use(recover.New())
		app.Use(logger.New())
	}
	route.Route(app)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
