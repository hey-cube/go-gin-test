package controllers

import (
	"go-gin-test/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListUsers(c *gin.Context) {
	users, err := models.ListUsers()

	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Abort()
	}

	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    users,
	})
}

func ShowUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Abort()
	}

	user, err := models.ShowUser(uint(id))

	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Abort()
	}

	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    user,
	})
}

func CreateUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)

	if err != nil {
		c.Status(http.StatusBadRequest)
		c.Abort()
	}

	if err := models.CreateUser(user); err != nil {
		c.Status(http.StatusInternalServerError)
		c.Abort()
	}

	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Abort()
	}

	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.Status(http.StatusBadRequest)
		c.Abort()
	}

	if err := models.UpdateUser(uint(id), user); err != nil {
		c.Status(http.StatusInternalServerError)
		c.Abort()
	}

	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Abort()
	}

	if err := models.DeleteUser(uint(id)); err != nil {
		c.Status(http.StatusInternalServerError)
		c.Abort()
	}

	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
	})
}
