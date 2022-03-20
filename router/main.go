package router

import (
	"bms/controller"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/middleware/logger"
	//"github.com/sirupsen/logrus"
	//"os"
)

var (
	bookController = controller.NewBookController()

	// logrus提供了New()函数来创建一个logrus的实例.
	// 项目中,可以创建任意数量的logrus实例.
	//log = logrus.New()
)

func Api() *iris.Application {
	app := iris.New()

	// 为当前logrus实例设置消息的输出,同样地,
	// 可以设置logrus实例的输出到任意io.writer
	//log.Out = os.Stdout

	// 为当前logrus实例设置消息输出格式为json格式.
	// 同样地,也可以单独为某个logrus实例设置日志级别和hook,这里不详细叙述.
	//log.Formatter = &logrus.JSONFormatter{}
	//
	//log.WithFields(logrus.Fields{
	//	"animal": "walrus",
	//	"size":   10,
	//}).Info("A group of walrus emerges from the ocean")

	//iris自带日志文件
	app.Logger().SetLevel("debug")
	//设置recover从panics恢复，设置log记录
	//app.Use(recover.New())
	app.Use(logger.New())
	SetupBook(app)

	return app
}

// SetupBook BOOK 相关接口
func SetupBook(party iris.Party) {
	party.PartyFunc("/book", func(book router.Party) {
		book.Use(controller.PreHandler)
		book.Get("/:id", bookController.Get)
		book.Post("/", bookController.Create)
		book.Post("/:id", bookController.Update)
		book.Delete("/:id", bookController.Delete)
	})
}
