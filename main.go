package main

import (
	"bms/src/server"
	"bms/src/server/utils"
)

func main() {
	utils.SetupConfig()
	utils.SetupDB()
	server.SetupServer()
}
