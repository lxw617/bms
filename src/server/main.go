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

	//加载html模板文件
	app.RegisterView(iris.HTML("./src/views", ".html"))
	app.Get("/", func(context iris.Context) {
		//绑定页面message 参数
		context.ViewData("message", "Hello world")
		//渲染模板文件
		context.View("hello.html")
	})

	//使用网络地址启动服务器
	//app.Run(iris.Addr(":3000"))

	//创建自定义或标准 net.Listener 并将其传递给 app.Run
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
