package controller

import (
	"bms/dao"
	"bms/model"
	"github.com/kataras/iris/v12"
)

// BookController ...
type BookController struct {
	dao *dao.BookDao
}

// NewBookController ...
func NewBookController() *BookController {
	return &BookController{dao: &dao.BookDao{}}
}

// Get ...
func (b *BookController) Get(ctx iris.Context) {
	bookID, _ := ctx.Params().GetInt("id")
	//fmt.Println(ctx.Values().Get("aaa"))
	book, _ := b.dao.GetBook(bookID)
	_, _ = ctx.JSON(book)
	ctx.Next()
}

// Create 创建书籍
func (b *BookController) Create(ctx iris.Context) {
	book := &model.Book{}
	_ = ctx.ReadJSON(book)
	_, _ = b.dao.CreateBook(book)
}

// Update ...
func (b *BookController) Update(ctx iris.Context) {
	book := &model.Book{}
	// TODO: 待办xxx
	// FIXME:
	// DEBUG:
	// OPTIMIZE:
	// HACK:
	// NOTE:
	_ = ctx.ReadJSON(book)
	book, _ = b.dao.UpdateBook(book)
	_, _ = ctx.JSON(book)
}

func (b *BookController) Delete(ctx iris.Context) {
	bookId, _ := ctx.Params().GetInt("id")
	_ = b.dao.DeleteBook(bookId)
	_, _ = ctx.JSON(map[string]interface{}{
		"status": "OK",
	})
}
