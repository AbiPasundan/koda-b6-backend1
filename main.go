package main

import (
	"net/http"
	_ "satu/docs"
	"satu/internal/handler"
	"satu/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func authMiddleware() gin.HandlerFunc {
	godotenv.Load()
	return func(ctx *gin.Context) {
		ctx.Header("Access-Controll-Allow-Origin", "localhost:5173")
		ctx.Header("Access-Controll-Allow-Headers", "content-type")
		if ctx.Request.Method == "OPTIONS" {
			ctx.Data(http.StatusOK, "", []byte(""))
		} else {
			ctx.Next()
		}
	}
}

// @BasePath godoc
//	@title			Base Path
//	@version		1.0.0
//	@description	This is minitask backend1
//	@termsOfService	http://swagger.io/terms/

//	@host						localhost:8881
//	@BasePath					/api/v1
//	@securityDefinitions.basic	BasicAuth

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	swaggo := r.Group("/api/v1")
	swaggo.Use(middleware.CookieTool())
	{
		swaggo.GET("/", handler.Home)
		swaggo.GET("/users/:id", handler.UserSearch)
		swaggo.DELETE("/users/:id", handler.Delete)
		swaggo.PUT("/users/:id", handler.Edit)
	}

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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8881")
	// r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
