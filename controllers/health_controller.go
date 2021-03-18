package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Healthcheck(c *gin.Context) {
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
	})
}
