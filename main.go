package main

import (
	"bms/server"
	"bms/utils"
)

func main() {
	utils.SetupDB()
	server.SetupServer()
}
