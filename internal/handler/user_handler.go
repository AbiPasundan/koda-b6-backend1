package handler

import (
	"fmt"
	"net/http"
	"net/mail"
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

// Home godoc
// @Summary      Get All Users
// @Description  Retrieve all users from the system
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200 {object} models.Response
// @Failure      500 {object} models.Response
// @Router       / [get]
func Home(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Data User",
		Results: ListUser,
	})
}

// UserSearch godoc
// @Summary      Get user by ID
// @Description  Retrieve a single user by its ID parameter
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Router       /users/{id} [get]
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

// Register godoc
// @Summary      Register Get
// @Description  Show Message in register
// @Tags         Register
// @Accept       json
// @Produce      json
// @Success      200 {object} models.Response
// @Failure      500 {object} models.Response
// @Router       /register [get]
func Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Halaman Register Silahkan isi Di dengan POST",
	})
}

// RegisterPost godoc
// @Summary Register Post
// @Description Register Process
// @Tags Register
// @Accept json
// @Produce json
// @Param        test body models.Users true "Login request body (email & password)"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /register [post]
func RegisterPost(ctx *gin.Context) {
	var data = models.Users{}
	var err = ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{
			Success: true,
			Message: "Something Gone Wrong",
		})
	} else {
		_, err := mail.ParseAddress(data.Email)
		if err != nil {
			ctx.JSON(400, models.Response{
				Success: false,
				Message: "That is not an email",
				Results: ListUser,
			})
			return
		} else {
			for x := range ListUser {
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
			fmt.Println(encoded)
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
		})
	}
}

// Login godoc
// @Summary      Login Get
// @Description  Show Message in Login
// @Tags         login
// @Accept       json
// @Produce      json
// @Success      200 {object} models.Response
// @Failure      500 {object} models.Response
// @Router       /login [get]
func Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Halaman Login Silahkan isi Di dengan POST",
	})
}

// LoginPost godoc
// @Summary Login Post
// @Description Login Process
// @Tags login
// @Accept json
// @Produce json
// @Param        request body models.Users true "Login request body (email & password)"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /login [post]
func LoginPost(ctx *gin.Context) {
	ctx.Header("Access-Controll-Allow-Origin", "localhost:5173")
	var data = models.Users{}

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid request body",
		})
		return
	}

	for _, user := range ListUser {
		if data.Email == user.Email {

			ok, err := argon2.VerifyEncoded(
				[]byte(data.Password),
				[]byte(user.Password),
			)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, models.Response{
					Success: false,
					Message: "Internal server error",
				})
				return
			}

			if ok {
				ctx.SetCookie("label", "ok", 3600, "/", "", false, true)

				ctx.JSON(http.StatusOK, models.Response{
					Success: true,
					Message: "Login success!",
				})
				return
			}
			break
		}
	}

	ctx.JSON(http.StatusBadRequest, models.Response{
		Success: false,
		Message: "Email or Password Wrong",
		Results: nil,
	})
}

// Delete godoc
// @Summary Delete func
// @Description Delete user by id
// @Tags delete
// @Accept json
// @Produce json
// @Param        id path int true "Delete User"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 404 {object} models.Response
// @Router /users/{id} [delete]
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

// Edit godoc
// @Summary      Update user
// @Description  Update user data by ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id path int true "User ID"
// @Param        request body models.UpdateUser true "Update user payload"
// @Success      200 {object} models.Response
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Router       /users/{id} [put]
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
