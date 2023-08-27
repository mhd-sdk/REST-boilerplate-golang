package main

import (
	"example.com/todo-app/database"
	"example.com/todo-app/docs"
	"example.com/todo-app/pkg/api_handlers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var RandomNumber int

// @title TODO API documentation
// @version 1.0
// @description API documentation for interacting with the TODO backend

// @contact.url
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api
// @schemes http
func main() {
	// init router
	router := gin.Default()
	router.SetTrustedProxies(nil)

	// init .env config
	viper.SetConfigFile("./../../.env")
	viper.ReadInConfig()
	port := viper.GetString("API_PORT")

	// init database connection
	database.Init()

	// init swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// routes
	v1 := router.Group("/api/v1")
	{
		todos := v1.Group("/todos")
		{
			todos.GET("/all", api_handlers.GetTodos)
			todos.POST("/new", api_handlers.CreateTodo)
			todos.GET("/:id", api_handlers.GetTodo)
			todos.PUT("/:id", api_handlers.EditTodo)
		}
	}
	router.Run("127.0.0.1:" + port)
}
