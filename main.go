package main

import (
	"go-gin-test/database"
	"go-gin-test/server"
)

func main() {
	database.Init()
	router := server.CreateRouter()
	router.Run(":3000")
}
