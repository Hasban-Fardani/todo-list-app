package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hasban-fardani/todo-list-app/pkg/configs"
)

var Mysql *sql.DB

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", configs.MysqlConnectionString)
	if err != nil {
		return nil, err
	}
	Mysql = db
	return db, nil
}

func init() {
	Connect()
}
