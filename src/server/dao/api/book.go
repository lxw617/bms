package api

import (
	"bms/src/server/model"
	"bms/src/server/utils"
)

type BookDao struct {
}

func (*BookDao) GetBook(bookId int) (*model.Book, error) {
	book := &model.Book{}
	if err := utils.BookStore_WR.Model(model.Book{}).Where("book_id = ?", bookId).Find(book).Error; err != nil {
		return nil, err
	}

	//sql := `SELECT * FROM t_book WHERE id = ?`
	//if err := utils.BookStore_WR.Raw(sql, bookID).Find(book).Error; err != nil {
	//	return nil, err
	//}

	return book, nil
}

func (*BookDao) CreateBook(book *model.Book) (*model.Book, error) {
	if err := utils.BookStore_WR.Create(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (*BookDao) UpdateBook(book *model.Book) (*model.Book, error) {
	err := utils.BookStore_WR.Model(model.Book{}).Where("book_id=?", book.ID).Updates(book).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (*BookDao) DeleteBook(bookId int) error {
	if err := utils.BookStore_WR.Where("id = ?", bookId).Delete(&model.Book{}).Error; err != nil {
		return err
	}
	return nil
}

func (*BookDao) GetAllBook() ([]model.Book, error) {
	var books []model.Book
	if err := utils.BookStore_WR.Model(model.Book{}).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
