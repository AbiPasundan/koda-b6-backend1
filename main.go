package main

import (
	"fmt"
	"net/http"
	"strconv"
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
	r := gin.Default()

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
			ctx.JSON(http.StatusOK, Response{
				Success: true,
				Message: "Something Gone Wrong",
			})
		} else {
			for x := range ListUser {
				if data.Email == ListUser[x].Email {
					ctx.JSON(400, Response{
						Success: true,
						Message: "Duplicated Email Not palid",
						Results: ListUser,
					})
					return
				}
			}
			ListUser = append(ListUser, Users{
				Id:       idCounter(),
				Email:    data.Email,
				Password: data.Password,
			})

			ctx.JSON(200, Response{
				Success: true,
				Message: "Back End is Running Well test",
				Results: ListUser,
			})
		}
	})

	r.GET("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		i, err := strconv.Atoi(id)
		// if i != 1 {
		// 	i--
		// }
		if i == 1 {
			i = 1
		}

		if err != nil {
			ctx.JSON(400, Response{
				Success: true,
				Message: "Bad Request",
				Results: 0,
			})
		}
		fmt.Println(i)

		for _, user := range ListUser {
			fmt.Println(user.Id)
			if int(user.Id) == i {
				ctx.JSON(200, Response{
					Success: true,
					Message: "berhasil",
					Results: ListUser[i-1],
				})
				return
			}
		}
		ctx.JSON(404, Response{
			Success: true,
			Message: "eubofceobu",
			Results: nil,
		})

	})

	r.Run("localhost:8888")
}
