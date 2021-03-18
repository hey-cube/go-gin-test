package middlewares

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CheckSession(c *gin.Context) {
	session := sessions.Default(c)
	// サンプルコードなのでセッションの判定は適当
	email := "foo@example.com"
	if session.Get("loginUser") != email {
		c.Status(http.StatusUnauthorized)
		c.Abort()
	} else {
		c.Next()
	}
}
