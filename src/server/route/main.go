package route

import (
	"bms/src/server/route/api"
	"github.com/kataras/iris/v12"
)

func Route(application *iris.Application) {
	api.SetupApi(application)
}
