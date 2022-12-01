package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hasban-fardani/todo-list-app/pkg/configs"
	"github.com/hasban-fardani/todo-list-app/pkg/middleware"
)

func main() {
	gin.SetMode("release")
	PORT := configs.BE_PORT

	server := gin.New()
	server.Use(gin.Logger(), gin.Recovery())

	api := server.Group("/api")

	api.GET("/", Index())
	api.POST("/login", Login())
	api.POST("/signup", Signup())

	usr := api.Group("/user")
	usr.Use(middleware.NeedLogin())
	usr.GET("/account", Index())
	usr.PUT("/account", EditUser())
	usr.DELETE("/account", DeleteUser())

	todo := api.Group("todo")
	todo.Use(middleware.NeedLogin())
	todo.GET("/", GetAllTodo())
	todo.POST("/add", AddTodo())

	todoGroup := api.Group("todo-group")
	todoGroup.Use(middleware.NeedLogin())
	todoGroup.GET("/", GetAllTodoGroup())
	todoGroup.POST("/add", AddTodoGroup())

	fmt.Println("server run on http://localhost:" + PORT)
	server.Run(":" + PORT)
}
