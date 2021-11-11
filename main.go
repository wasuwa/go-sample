package main

import (
	"twitter-app/config"
	"twitter-app/database"
	"twitter-app/server"
)

func main() {
	config.Init("config/environments/", "development")
	database.Init("db.url")
	defer database.Close()
	server.Init()
}
