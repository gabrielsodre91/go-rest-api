package routes

import (
	"github.com/gabrielsodre91/api-gin/auth"
	"github.com/gabrielsodre91/api-gin/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		auths := main.Group("auth")
		{
			auths.POST("/login", auth.Login)
		}
		books := main.Group("books", auth.Guard())
		{
			books.GET("/", controllers.GetAllBooks)
			books.GET("/:id", controllers.GetBook)
			books.PUT("/", controllers.UpdateBook)
			books.POST("/", controllers.CreateBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
	}

	return router
}