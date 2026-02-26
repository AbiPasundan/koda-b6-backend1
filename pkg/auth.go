package pkg

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
)

var data = Users{}

func Auth() {
	r := gin.Default()

	r.GET("/register", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Message: "Halaman Register Silahkan isi Di dengan POST",
		})
	})

	r.POST("/register", func(ctx *gin.Context) {
		var err = ctx.ShouldBindJSON(&data)
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

	r.GET("/login", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Message: "Halaman Login Silahkan isi Di dengan POST",
		})
	})

	r.POST("/login", func(ctx *gin.Context) {
		// ctx.JSON(http.StatusOK, Response{
		// 	Success: true,
		// 	Message: "Something Gone Wrong",
		// })
		for _, user := range ListUser {

			// argon := argon2.DefaultConfig()

			// encoded, err := argon.HashEncoded([]byte(data.Password))

			// ok, err := argon2.VerifyEncoded([]byte(user.Password), encoded)
			ok, err := argon2.VerifyEncoded(
				[]byte(data.Password),
				[]byte(user.Password),
			)

			if err != nil {
				panic(err)
			}

			fmt.Println(ok)

			if data.Email == user.Email && ok {
				ctx.JSON(http.StatusOK, Response{
					Success: true,
					Message: "Ok email dan password sama",
				})
			} else {
				ctx.JSON(400, Response{
					Success: true,
					Message: "Email or Password Wrong",
					Results: nil,
				})
			}
		}
	})

	r.Run("localhost:8888")
}
