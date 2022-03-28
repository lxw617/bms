package api

import (
	"bms/src/server/route/api/book"
	"github.com/kataras/iris/v12"
)

func SetupApi(application *iris.Application) {

	api := application.Party("/api")
	book.SetUp(api)
}
