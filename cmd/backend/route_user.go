package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/hasban-fardani/todo-list-app/pkg/auth"
	"github.com/hasban-fardani/todo-list-app/pkg/configs"
	"github.com/hasban-fardani/todo-list-app/pkg/database"
	"github.com/hasban-fardani/todo-list-app/pkg/models"
)

var db = database.Mysql

type bodyStruct struct {
	From models.User `json:"From"`
	To   models.User `json:"to"`
}

func Index() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		status := http.StatusOK
		message := "you're logged in"
		token := ctx.GetHeader("Authorization")
		tokenStr := strings.Replace(token, "Bearer ", "", -1)

		is_login, err := auth.LoginWithToken(tokenStr)
		if err != nil {
			fmt.Println("token: " + tokenStr)
			fmt.Println(err.Error())
			message = "you're not login"
		}

		ctx.JSON(status, gin.H{
			"status":   status,
			"is_login": is_login,
			"message":  message,
		})
	}
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var ok bool
		var err error
		var userData models.User

		var message = "login success"
		var status = http.StatusAccepted

		tokenStr := ctx.GetHeader("Authorization")

		err = ctx.BindJSON(&userData)
		if err != nil {
			status = http.StatusInternalServerError
			message = err.Error()
		}

		if tokenStr != "" {
			ok, err = auth.LoginWithToken(tokenStr)
			if err != nil || !ok {
				status = http.StatusInternalServerError
				message = err.Error()
			}
		} else {
			user, ok, err := auth.LoginWithData(userData)
			if err != nil || !ok {
				ctx.JSON(http.StatusInternalServerError,
					gin.H{
						"message": "failed to login, please check whether the data you entered is correct",
					},
				)
				return
			}

			claims := models.UserClaims{
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: time.Now().Add(configs.JwtTimeOut).Unix(),
					Issuer:    configs.AppName,
				},
				User: user,
			}
			tokenStr, err = auth.CreateUserToken(claims)
			if err != nil {
				status = http.StatusInternalServerError
				message = err.Error()
			}
			ctx.Header("authorization", "Bearer "+tokenStr)
		}

		ctx.JSON(status, gin.H{
			"status":  status,
			"message": message,
			"token":   tokenStr,
		})
	}
}

func Signup() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userData models.User

		var status = http.StatusAccepted
		var message = "register succes"

		if err := ctx.BindJSON(&userData); err != nil {
			status = http.StatusInternalServerError
			message = err.Error()
		}
		ok, err := auth.Signup(userData)
		if err != nil || !ok {
			status = http.StatusInternalServerError
			message = "user telah terdaftar"
		}

		//
		if err != nil {
			status = http.StatusInternalServerError
			message = err.Error()
		}

		ctx.JSON(status, gin.H{
			"status":  status,
			"message": message,
			"time":    time.Now(),
		})
	}
}

func EditUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req bodyStruct
		var err error

		var null_usr = models.User{}
		var status = http.StatusOK
		var message = "success update data"

		if err := ctx.BindJSON(&req); err != nil {
			fmt.Println(err.Error())
		}

		// jika req.From tidak memiliki value
		if req.From == null_usr {
			fmt.Println(req.From.NamaLengkap)
			fmt.Println(req.To.NamaLengkap)
			message = "no user selected"
			status = http.StatusInternalServerError
		}

		if req.To.Username != "" {
			db.Exec(
				"update user set username = ? where username = ?",
				req.To.Username, req.From.Username)
		}
		if req.To.NamaLengkap != "" {
			db.Exec(
				"update user set namaLengkap = ? where namaLengkap = ?",
				req.To.NamaLengkap, req.From.NamaLengkap)
		}
		if req.To.Email != "" {
			db.Exec(
				"update user set email = ? where email = ?",
				req.To.Email, req.From.Email)
		}
		if req.To.Password != "" {
			db.Exec(
				"update user set password = ? where password = ?",
				req.To.Password, req.From.Password)
		}

		db.Exec(
			"update user set lastEditeAt = CURRENT_TIMESTAMP where username = ?",
			req.From.Username)

		if err != nil {
			status = http.StatusInternalServerError
			message = err.Error()
		}
		ctx.JSON(status, gin.H{
			"status":  status,
			"message": message,
		})
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var usrInDb models.User
		var err error

		var status = http.StatusAccepted
		var message = "success delete user"

		userData, exists := ctx.Get("user")
		if !exists {
			status = http.StatusInternalServerError
			message = "you're not login!"
		}
		user := userData.(models.User)
		err = db.QueryRow(
			"select * from todo where idUser = ?", user.Id,
		).Scan(&usrInDb)
		if usrInDb.Username == "" {
			status = http.StatusBadRequest
			message = "user not exists"
		}
		_, err = db.Query("delete from todo where idUser = ?", user.Id)
		_, err = db.Query("delete from todoGroup where idUser = ?", user.Id)
		_, err = db.Query("delete from user where id = ?", user.Id)

		ctx.Header("Authorization", "")
		if err != nil {
			status = http.StatusInternalServerError
			message = err.Error()
		}
		ctx.JSON(status, gin.H{
			"status":  status,
			"message": message,
		})
	}
}
