package server

import (
	"go-gin-test/controllers"
	"go-gin-test/middlewares"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	v1 := router.Group("/api/v1")
	{
		v1.GET("/healthcheck", controllers.Healthcheck)
		v1.POST("/login", controllers.Login)

		auth := v1.Group("/auth")
		auth.Use(middlewares.CheckSession)
		{
			users := auth.Group("/users")
			{
				users.GET("", controllers.ListUsers)
				users.GET("/:id", controllers.ShowUser)
				users.POST("", controllers.CreateUser)
				users.PATCH("/:id", controllers.UpdateUser)
				users.DELETE("/:id", controllers.DeleteUser)
			}
		}
	}

	return router
}
