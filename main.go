package main

import (
	"twitter-app/config"
	"twitter-app/database"
	"twitter-app/server"
)

func main() {
	config.Init("development")
	server.Init()
	defer database.Close()
}
