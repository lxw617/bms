package router

import (
	"bms/controller"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

var (
	bookController = controller.NewBookController()
)

func Api() *iris.Application {
	app := iris.New()

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
