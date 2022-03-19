package dao

import (
	"bms/model"
	"bms/utils"
)

type BookDao struct {
}

func (*BookDao) GetBook(bookId int) (*model.Book, error) {
	book := &model.Book{}
	if err := utils.BookDB.Model(model.Book{}).Where("id = ?", bookId).Find(book).Error; err != nil {
		return nil, err
	}

	//sql := `SELECT * FROM t_book WHERE id = ?`
	//if err := utils.BookDB.Raw(sql, bookID).Find(book).Error; err != nil {
	//	return nil, err
	//}

	return book, nil
}

func (*BookDao) CreateBook(book *model.Book) (*model.Book, error) {
	if err := utils.BookDB.Create(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (*BookDao) UpdateBook(book *model.Book) (*model.Book, error) {
	err := utils.BookDB.Model(model.Book{}).Where("id=?", book.ID).Updates(book).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (*BookDao) DeleteBook(bookId int) error {
	if err := utils.BookDB.Where("id = ?", bookId).Delete(&model.Book{}).Error; err != nil {
		return err
	}
	return nil
}
