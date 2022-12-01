package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hasban-fardani/todo-list-app/pkg/models"
)

func GetAllTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var result []models.Todo
		var each models.Todo

		var status = http.StatusAccepted
		var message = ""

		user, _ := ctx.Get("user")
		rows, err := db.Query("select * from todo where idUser = ?", user.(models.User).Id)
		if err != nil {
			status = http.StatusInternalServerError
			message = err.Error()
			fmt.Println(message)
		}
		for rows.Next() {
			err = rows.Scan(
				&each.Id, &each.IdUser, &each.IdGroup,
				&each.Nama, &each.Deskripsi,
				&each.StartAt, &each.EndAt,
			)
			if err != nil || rows.Err() != nil {
				status = http.StatusInternalServerError
				message = err.Error()
			}
			result = append(result, each)
		}

		ctx.JSON(status, gin.H{
			"status":  status,
			"message": message,
			"data":    result,
		})
	}
}

func GetAllTodoGroup() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var result []models.TodoGroup
		var each models.TodoGroup

		var status = http.StatusAccepted
		var message = ""

		user, _ := ctx.Get("user")
		rows, err := db.Query("select * from todoGroup where idUser = ?", user.(models.User).Id)
		if err != nil {
			status = http.StatusInternalServerError
			message = err.Error()
		}
		for rows.Next() {
			err = rows.Scan(
				&each.Id, &each.Nama,
				&each.SkalaPrioritas,
				&each.Warna, &each.IdUser,
			)
			if err != nil || rows.Err() != nil {
				status = http.StatusInternalServerError
				message = err.Error()
			}
			result = append(result, each)
			fmt.Println(each)
		}

		ctx.JSON(status, gin.H{
			"status":  status,
			"message": message,
			"data":    result,
		})
	}
}

func AddTodoGroup() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requests models.TodoGroup

		var message = "success add new todo group"
		if err := ctx.BindJSON(&requests); err != nil {
			return
		}

		_, err := db.Exec(
			"insert into todoGroup (nama, skalaPrioritas, warna, idUser) values (?, ?, ?, ?)",
			requests.Nama, requests.SkalaPrioritas, requests.Warna, requests.IdUser,
		)

		if err != nil {
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	}
}

func AddTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request models.Todo

		var message = "success add new todo"
		if err := ctx.BindJSON(&request); err != nil {
			fmt.Println(err.Error())
			return
		}

		// request.StartAt, _ = time.Parse("", request.StartAt.String())
		userData, _ := ctx.Get("user")
		user := userData.(models.User)
		_, err := db.Exec(
			"insert into todo (idUser, idGroup, nama, deskripsi, startAt, endAt) values (?, ?, ?, ?, ?, ?)",
			user.Id, request.IdGroup, request.Nama, request.Deskripsi, request.StartAt, request.EndAt,
		)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	}
}

type bodyStructTodo struct {
	From models.Todo `json:"from"`
	To   models.Todo `json:"to"`
}

type bodyStructTodoGroup struct {
	From models.Todo `json:"from"`
	To   models.Todo `json:"to"`
}

func EditTodoGroup() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func EditTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body bodyStructTodo
		var err error

		if err := ctx.BindJSON(&body); err != nil {

			return
		}
		if body.To.Deskripsi != "" {
			_, err = db.Exec(
				"update todoGroup set deskripsi = ? where id = ?",
				body.To.Deskripsi, body.From.Id,
			)
		}
		if body.To.Nama != "" {
			_, err = db.Exec(
				"update todoGroup set nama = ? where id = ?",
				body.To.Nama, body.From.Id,
			)
		}

		// if body.To.StartAt.String() != "" {
		// 	_, err = db.Exec(
		// 		// "update todoGroup set nama = ? where id = ?",
		// 		body.To.Nama, body.From.Id,
		// 	)
		// }
		if err != nil {
			log.Fatalln(err.Error())
		}
	}
}
