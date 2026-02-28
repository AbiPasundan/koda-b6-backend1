package handler

import (
	"fmt"
	"net/http"
	"satu/internal/models"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
)

var ListUser []models.Users
var Counter int64

func idCounter() int64 {
	return atomic.AddInt64(&Counter, 1)
}

func Home(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Data User",
		Results: ListUser,
	})
}

func UserSearch(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(400, models.Response{
			Success: false,
			Message: "Bad Request",
			Results: nil,
		})
		return
	}

	for _, user := range ListUser {
		fmt.Println(user.Id)
		if int(user.Id) == i {
			ctx.JSON(200, models.Response{
				Success: true,
				Message: "berhasil",
				Results: user,
			})
			return
		}
	}

	ctx.JSON(404, models.Response{
		Success: false,
		Message: "User not found",
		Results: nil,
	})
}

func Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Halaman Register Silahkan isi Di dengan POST",
	})
}

func RegisterPost(ctx *gin.Context) {
	var data = models.Users{}
	var err = ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{
			Success: true,
			Message: "Something Gone Wrong",
		})
	} else {
		for x := range ListUser {
			wordToCheck := "@"

			if !strings.Contains(data.Email, wordToCheck) {
				ctx.JSON(400, models.Response{
					Success: false,
					Message: "That is not an email",
					Results: ListUser,
				})
				return
			} else {
				if data.Email == ListUser[x].Email {
					ctx.JSON(400, models.Response{
						Success: false,
						Message: "Duplicated Email Not palid",
						Results: ListUser,
					})
					return
				}
				if len(data.Password) <= 8 {
					ctx.JSON(400, models.Response{
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

		ListUser = append(ListUser, models.Users{
			Id:       idCounter(),
			Email:    data.Email,
			Name:     data.Name,
			Password: string(encoded),
		})

		ctx.JSON(200, models.Response{
			Success: true,
			Message: "Berhasil register",
			Results: ListUser,
		})
	}
}

func Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Halaman Login Silahkan isi Di dengan POST",
	})
}
func LoginPost(ctx *gin.Context) {
	var data = models.Users{}
	for _, user := range ListUser {
		ok, err := argon2.VerifyEncoded(
			[]byte(data.Password),
			[]byte(user.Password),
		)

		if err != nil {
			panic(err)
		}

		if data.Email == user.Email && ok {
			ctx.SetCookie("label", "ok", 100, "/", "localhost", false, true)

			ctx.String(200, "Login success!")
			return
		}
	}
	ctx.ShouldBindJSON(&data)
	ctx.JSON(400, models.Response{
		Success: false,
		Message: "Email or Password Wrong",
		Results: nil,
	})
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	i, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid ID",
			Results: nil,
		})
		return
	}

	for index, user := range ListUser {
		if int(user.Id) == i {

			ListUser = append(ListUser[:index], ListUser[index+1:]...)

			ctx.JSON(http.StatusOK, models.Response{
				Success: true,
				Message: "User berhasil dihapus",
				Results: ListUser,
			})
			ctx.SetCookie("gin_cookie", "", -1, "/", "localhost", false, true)
			ctx.String(http.StatusOK, "Cookie removed")
			return
		}
	}

	ctx.JSON(http.StatusNotFound, models.Response{
		Success: false,
		Message: "User tidak ditemukan",
		Results: nil,
	})
}

func Edit(ctx *gin.Context) {
	id := ctx.Param("id")

	i, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid ID",
			Results: nil,
		})
		return
	}

	var input models.UpdateUser
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid JSON",
			Results: nil,
		})
		return
	}

	for index, user := range ListUser {
		if int(user.Id) == i {

			if input.Name != "" {
				ListUser[index].Name = input.Name
			}

			if input.Email != "" {
				if !strings.Contains(input.Email, "@") {
					ctx.JSON(http.StatusBadRequest, models.Response{
						Success: false,
						Message: "Email tidak valid",
						Results: nil,
					})
					return
				}
				ListUser[index].Email = input.Email
			}

			if input.Password != "" {
				if len(input.Password) <= 8 {
					ctx.JSON(http.StatusBadRequest, models.Response{
						Success: false,
						Message: "Password terlalu lemah",
						Results: nil,
					})
					return
				}

				argon := argon2.DefaultConfig()
				encoded, err := argon.HashEncoded([]byte(input.Password))
				if err != nil {
					panic(err)
				}

				ListUser[index].Password = string(encoded)
			}

			ctx.JSON(http.StatusOK, models.Response{
				Success: true,
				Message: "User berhasil diupdate",
				Results: ListUser[index],
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, models.Response{
		Success: false,
		Message: "User tidak ditemukan",
		Results: nil,
	})
}
