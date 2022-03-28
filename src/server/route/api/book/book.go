package book

import (
	bookCtrls "bms/src/server/controller/api/book"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

func SetUp(party iris.Party) {
	var (
		bookController = &bookCtrls.BookController{}
		//bookController = bookCtrls.NewBookController()
	)
	party.PartyFunc("/book", func(book router.Party) {
		//book.Use(controller.PreHandler)
		book.Get("/", bookController.GetAll)
		book.Get("/:id", bookController.Get)
		book.Post("/", bookController.Create)
		book.Post("/:id", bookController.Update)
		book.Delete("/:id", bookController.Delete)
		book.Put("/batchCreate", bookController.BatchCreate)
	})
}
