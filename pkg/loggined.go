package pkg

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"Message"`
	Results any    `json:"Results"`
}
type Users struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"Password"`
}

var ListUser []Users

var Counter int64

func idCounter() int64 {
	return atomic.AddInt64(&Counter, 1)
}

func Loggined() {
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
				wordToCheck := "@"

				if !strings.Contains(data.Email, wordToCheck) {
					ctx.JSON(400, Response{
						Success: false,
						Message: "That is not an email",
						Results: ListUser,
					})
					return
				} else {
					if data.Email == ListUser[x].Email {
						ctx.JSON(400, Response{
							Success: false,
							Message: "Duplicated Email Not palid",
							Results: ListUser,
						})
						return
					}
					if len(data.Password) <= 8 {
						ctx.JSON(400, Response{
							Success: true,
							Message: "Password terlalu lemah",
							Results: ListUser,
						})
						return
					}
				}
			}

			argon := argon2.DefaultConfig()
			encoded, err := argon.HashEncoded([]byte(data.Password))
			if err != nil {
				panic(err)
			}

			ListUser = append(ListUser, Users{
				Id:       idCounter(),
				Email:    data.Email,
				Name:     data.Name,
				Password: string(encoded),
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
