package main

import (
	"main/database"
	"main/environment"
	"main/logging"
	"main/server"
)

func main() {

	logging.Init()

	environment.InitEnv()

	database.Init()

	server.Start()

}
