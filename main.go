package main

import (
	"satu/pkg"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// pkg.OriginAuth()
	authorized := r.Group("/")
	authorized.Use(pkg.CookieTool())
	{
		authorized.GET("/", pkg.Home)
		authorized.GET("/user/:id", pkg.UserSearch)
	}
	// r.GET("/", pkg.CookieTool(), pkg.Home)
	// r.GET("/users/:id", pkg.CookieTool(), pkg.UserSearch)
	r.GET("/register", pkg.Register)
	r.POST("/register", pkg.RegisterPost)
	r.GET("/login", pkg.Login)
	r.POST("/login", pkg.LoginPost)
	r.Run("localhost:8888")
}
