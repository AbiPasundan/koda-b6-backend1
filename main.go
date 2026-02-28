package main

import (
	"satu/internal/handler"
	"satu/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	authorized := r.Group("/")
	authorized.Use(middleware.CookieTool())
	{
		authorized.GET("/", handler.Home)
		authorized.GET("/users/:id", handler.UserSearch)
		authorized.DELETE("/users/:id", handler.Delete)
		authorized.PUT("/users/:id", handler.Edit)
	}
	r.GET("/register", handler.Register)
	r.GET("/login", handler.Login)
	r.POST("/register", handler.RegisterPost)
	r.POST("/login", handler.LoginPost)
	r.Run(":8888")
}
