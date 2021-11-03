package main

import (
	"twitter-app/config"
	"twitter-app/server"
)

func main() {
	config.Init("development")
	server.Init()
}
