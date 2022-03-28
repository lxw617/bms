package utils

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hoisie/mustache"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func _SetupDB(name string, readOnly bool) *gorm.DB {
	isDevelopment := viper.GetBool("isDevelopment")
	instance := "wr"
	if readOnly {
		instance = "r"
	}

	prefix := mustache.Render("resource.database.{{name}}.{{instance}}", map[string]interface{}{
		"name":     name,
		"instance": instance,
	})
	dialect := viper.GetString(prefix + ".dialect")
	user := viper.GetString(prefix + ".user")
	password := viper.GetString(prefix + ".password")
	database := viper.GetString(prefix + ".database")
	host := viper.GetString(prefix + ".host")
	port := viper.GetString(prefix + ".port")
	//maxIdle := viper.GetInt(prefix + ".max-idle")
	//maxOpen := viper.GetInt(prefix + ".max-open")
	url := mustache.Render("{{user}}:{{password}}@tcp({{host}}:{{port}})/{{database}}?charset=utf8&parseTime=True&loc=Local", map[string]interface{}{
		"user":     user,
		"password": password,
		"database": database,
		"host":     host,
		"port":     port,
	})
	//url := mustache.Render("{{user}}:{{password}}@/{{database}}?charset=utf8&parseTime=True&loc=Local", map[string]interface{}{
	//	"user":     user,
	//	"password": password,
	//	"database": database,
	//})
	db, err := gorm.Open(dialect, url)
	//db, err := gorm.Open("mysql", "root:root@/bookstore?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		//Logger.Warningln("failed to connect database:", database, err.Error())
		fmt.Println(err)
		panic("failed to connect database")
	}
	db.LogMode(isDevelopment)
	//db.DB().SetMaxIdleConns(maxIdle)
	//db.DB().SetMaxOpenConns(maxOpen)
	return db
}

func SetupDB() {
	BookStore_WR = _SetupDB("bookstore", false)
	BookStore_R = _SetupDB("bookstore", true)
}

var (
	BookStore_WR *gorm.DB
	BookStore_R  *gorm.DB
)
