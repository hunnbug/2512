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

	environment.InitS3Enviroment()

	database.Init()

	server.Start()

}
