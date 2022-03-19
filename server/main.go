package server

import (
	"bms/router"
)

func SetupServer() {
	app := router.Api()

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
