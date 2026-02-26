package main

import (
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"Message"`
	Results any    `json:"Results"`
}
type Users struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"Password"`
}

var ListUser []Users

var counter int64

func idCounter() int64 {
	return atomic.AddInt64(&counter, 1)
}

func main() {
	fmt.Println(idCounter())
	r := gin.Default()

	// r.GET("/", func(ctx *gin.Context) {
	// 	ctx.Data(200, "text/plain", []byte("hello"))
	// 	ctx.JSON(200, Response{
	// 		Success: true,
	// 		Message: "Back End is Running Well test",
	// 		Results: ListUser,
	// 	})
	// })

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Message: "Data User",
			Results: ListUser,
		})
	})

	r.POST("/", func(ctx *gin.Context) {
		data := Users{}
		err := ctx.ShouldBindJSON(&data)

		if err != nil {
			// ctx.JSON(400, Response{
			// 	Success: true,
			// 	Message: "Back End is Running Well test",
			// })
			ctx.JSON(http.StatusOK, Response{
				Success: true,
				Message: "Something Gone Wrong",
			})
		} else {
			ListUser = append(ListUser, Users{
				Id:       idCounter(),
				Email:    data.Email,
				Password: data.Password,
			})
			ctx.Data(200, "text/plain", []byte("hello"))
			ctx.JSON(200, Response{
				Success: true,
				Message: "Back End is Running Well test",
				Results: ListUser,
			})
		}
	})

	r.GET("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		if id == "5" {
			ctx.JSON(200, Response{
				Success: true,
				Message: fmt.Sprintf("your id is %s", id),
			})
		} else {
			ctx.JSON(200, Response{
				Success: true,
				Message: fmt.Sprintf("Saha sia %s", id),
			})
		}

	})

	r.Run("localhost:8888")
}

// package main

// import (
// 	"fmt"
// 	"net/http" // Recommended to use net/http constants

// 	"github.com/gin-gonic/gin"
// )

// type Response struct {
// 	Success bool   `json:"success"`
// 	Message string `json:"message"` // Changed to lowercase 'm' for convention
// 	Results any    `json:"results"` // Changed to lowercase 'r' for convention
// }

// type Users struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"` // Changed to lowercase 'p' for convention
// }

// var ListUser []Users

// func main() {
// 	fmt.Println(123)
// 	r := gin.Default()

// 	r.GET("/", func(ctx *gin.Context) {
// 		ctx.JSON(http.StatusOK, Response{
// 			Success: true,
// 			Message: "Back End is Running Well test",
// 			Results: ListUser,
// 		})
// 	})

// 	r.POST("/", func(ctx *gin.Context) {
// 		data := Users{}
// 		err := ctx.ShouldBindJSON(&data)

// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, Response{
// 				Success: false,
// 				Message: "Invalid request body",
// 			})
// 		} else {
// 			ListUser = append(ListUser, data)
// 			ctx.JSON(http.StatusOK, Response{
// 				Success: true,
// 				Message: "User added successfully",
// 				Results: ListUser,
// 			})
// 		}
// 	})

// 	r.GET("/users/:id", func(ctx *gin.Context) {
// 		id := ctx.Param("id")

// 		if id == "5" {
// 			ctx.JSON(http.StatusOK, Response{
// 				Success: true,
// 				Message: fmt.Sprintf("your id is %s", id),
// 			})
// 		} else {
// 			ctx.JSON(http.StatusOK, Response{
// 				Success: true,
// 				Message: fmt.Sprintf("Saha sia %s", id),
// 			})
// 		}
// 	})

// 	r.Run("localhost:8888")
// }
