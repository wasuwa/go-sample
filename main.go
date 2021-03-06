package main

import (
	"twitter-app/config"
	"twitter-app/database"
	"twitter-app/server"
)

func main() {
	config.Init("development")
	database.Init()
	defer database.Close()
	server.Init()
}
