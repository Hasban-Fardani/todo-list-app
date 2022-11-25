package main

import (
	"fmt"
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
				&each.IdUser, &each.NamaKegiatan, &each.Deskripsi,
				&each.StartAt, &each.EndAt, &each.IdGroup,
				&each.Id,
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
				&each.Id,
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

func AddTodoGroup() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requests models.TodoGroup

		var message = "success add new todo group"
		if err := ctx.BindJSON(&requests); err != nil {
			return
		}

		_, err := db.Query(
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
			return
		}

		_, err := db.Query(
			"insert into todo (idUser, idGroup, namaKegiatan, deskripsi, startAt, endAt) values (?, ?, ?, ?, ?, ?)",
			request.IdUser, request.IdGroup, request.NamaKegiatan, request.Deskripsi, request.StartAt, request.EndAt,
		)

		if err != nil {
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	}
}

func EditTodoGroup() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func EditTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
