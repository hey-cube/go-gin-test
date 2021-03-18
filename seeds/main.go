package main

import (
	"fmt"
	"go-gin-test/database"
	"go-gin-test/models"
	"strconv"
)

func main() {
	database.Init()
	db := database.GetDB()

	for i := 0; i < 3; i++ {
		user := models.User{Name: "hey_cube" + strconv.Itoa(i)}
		if err := db.Create(&user).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}
}
