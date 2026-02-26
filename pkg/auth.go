package pkg

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
)

var data = Users{}

func CookieTool() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie("label"); err == nil {
			if cookie == "ok" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden with no cookie"})
		c.Abort()
	}
}

func Home(ctx *gin.Context) {
	// r.GET("/", CookieTool(), func(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Data User",
		Results: ListUser,
	})
	// })
}

func UserSearch(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(400, Response{
			Success: false,
			Message: "Bad Request",
			Results: nil,
		})
		return
	}

	for _, user := range ListUser {
		fmt.Println(user.Id)
		if int(user.Id) == i {
			ctx.JSON(200, Response{
				Success: true,
				Message: "berhasil",
				Results: user,
			})
			return
		}
	}

	ctx.JSON(404, Response{
		Success: false,
		Message: "User not found",
		Results: nil,
	})
}

func Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Halaman Register Silahkan isi Di dengan POST",
	})
}

func RegisterPost(ctx *gin.Context) {
	// r.POST("/register", func(ctx *gin.Context) {
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
}

func Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Halaman Login Silahkan isi Di dengan POST",
	})
}
func LoginPost(ctx *gin.Context) {
	for _, user := range ListUser {
		ok, err := argon2.VerifyEncoded(
			[]byte(data.Password),
			[]byte(user.Password),
		)

		if err != nil {
			panic(err)
		}

		if data.Email == user.Email && ok {
			ctx.SetCookie("label", "ok", 30, "/", "localhost", false, true)

			ctx.String(200, "Login success!")
			return
		}
	}
	ctx.ShouldBindJSON(&data)
	ctx.JSON(400, Response{
		Success: false,
		Message: "Email or Password Wrong",
		Results: nil,
	})
}
