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
		r.DELETE("/users/:id", handler.Delete)
		r.PUT("/users/:id", handler.Edit)
	}
	r.GET("/register", handler.Register)
	r.POST("/register", handler.RegisterPost)
	r.GET("/login", handler.Login)
	r.POST("/login", handler.LoginPost)
	r.Run("localhost:8888")
}
