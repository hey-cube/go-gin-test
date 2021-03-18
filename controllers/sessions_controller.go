package controllers

import (
	"go-gin-test/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var request models.LoginRequest
	err := c.BindJSON(&request)

	if err != nil {
		c.Status(http.StatusBadRequest)
		c.Abort()
	}

	// サンプルコードなのでログイン処理は適当
	email := "foo@example.com"
	password := "password1234"

	if request.Email != email || request.Password != password {
		c.Status(http.StatusBadRequest)
		c.Abort()
	}

	session := sessions.Default(c)
	session.Set("loginUser", email)
	session.Save()
	c.Status(http.StatusOK)
}
