package main

import (
	"main/database"
	"main/environment"
	"main/server"
)

func main() {

	environment.InitEnv()

	database.Init()

	server.Start()

}
