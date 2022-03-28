package model

type Book struct {
	ID      int
	Name    string
	Author  string
	Price   float64
	Sales   int
	Stock   int
	ImgPath string
}

func (*Book) TableName() string {
	return "t_book"
}
