package main

import (
	"fmt"
	"net/http"
	"os"
	"satu/internal/handler"
	"satu/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

func main() {
	r := gin.Default()

	// r.Use(authMiddleware())

	// r.OPTIONS("/login", func(ctx *gin.Context) {
	// 	ctx.Header("Access-Controll-Allow-Origin", "localhost:5173")
	// 	ctx.Data(http.StatusOK, "", []byte(""))
	// })

	r.Use(cors.Default())

	// ConnConfig, _ := pgx.ParseConfig("")

	// conn, err := pgx.Connect(context.Background(), ConnConfig.ConnString())

	// if err != nil {
	// 	fmt.Println("connection is failed")
	// 	fmt.Println(err)
	// }

	// r.GET("", func(ctx *gin.Context) {
	// 	rows, err := conn.Query(context.Background(),
	// 		`SELECT * FROM users`,
	// 	)

	// 	users, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[Users]())

	// 	if err != nil {
	// 		ctx.JSON(400, models.Response{
	// 			Success: false,
	// 			Message: "Bad Request",
	// 			Results: nil,
	// 		})
	// 		return
	// 	}

	// 	ctx.JSON(http.StatusOK, models.Response{
	// 		Success: true,
	// 		Message: "Data User",
	// 		Results: users,
	// 	})
	// })

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

	// r.Run(":8888")
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
