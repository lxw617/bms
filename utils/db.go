package utils

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var BookDB *gorm.DB

func SetupDB() {
	db, err := gorm.Open("mysql", "root:root@/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("connect database error")
	}
	BookDB = db
	//defer func() {
	//	_ = db.Close()
	//}()
}
