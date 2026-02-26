package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() {
	r := gin.Default()

	r.GET("/register", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Message: "Halaman Register Silahkan isi Di dengan POST",
		})
	})

	r.POST("/register", func(ctx *gin.Context) {
	})

	r.GET("/login", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Message: "Halaman Login Silahkan isi Di dengan POST",
		})
	})

	r.POST("/login", func(ctx *gin.Context) {
	})

	r.Run("localhost:8888")
}
