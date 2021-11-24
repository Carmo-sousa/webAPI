package routes

import (
	"github.com/Carmo-sousa/webAPI/controllers"
	"github.com/Carmo-sousa/webAPI/server/middlewares"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		main.POST("login", controllers.Login)

		users := main.Group("user")
		{
			users.GET("/", controllers.ShowAllUsers)
			users.POST("/", controllers.CreateUser)
			users.GET("/:id", controllers.ShowUser)
		}

		books := main.Group("books", middlewares.Auth())
		{
			books.GET("/", controllers.ShowAllBooks)
			books.POST("/", controllers.CreateBook)
			books.GET("/:id", controllers.ShowBook)
			books.PUT("/:id", controllers.UpdateBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
	}

	return router
}
