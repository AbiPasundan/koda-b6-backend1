package main

import (
	"satu/pkg"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	authorized := r.Group("/")
	authorized.Use(pkg.CookieTool())
	{
		authorized.GET("/", pkg.Home)
		authorized.GET("/users/:id", pkg.UserSearch)
	}
	r.GET("/register", pkg.Register)
	r.POST("/register", pkg.RegisterPost)
	r.GET("/login", pkg.Login)
	r.POST("/login", pkg.LoginPost)
	r.Run("localhost:8888")
}
