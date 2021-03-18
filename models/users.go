package models

import (
	"go-gin-test/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `binding:"required"`
}

func ListUsers() ([]User, error) {
	db := database.GetDB()
	var users []User
	result := db.Find(&users)

	return users, result.Error
}

func ShowUser(id uint) (User, error) {
	db := database.GetDB()
	var user User
	result := db.Find(&user, id)

	return user, result.Error
}

func CreateUser(user User) error {
	db := database.GetDB()
	result := db.Create(&user)

	return result.Error
}

func UpdateUser(id uint, user User) error {
	db := database.GetDB()
	result := db.Model(User{Model: gorm.Model{ID: id}}).Updates(user)

	return result.Error
}

func DeleteUser(id uint) error {
	db := database.GetDB()
	result := db.Delete(&User{}, id)

	return result.Error
}
