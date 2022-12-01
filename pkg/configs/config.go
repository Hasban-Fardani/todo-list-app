package configs

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	MysqlConnectionString = os.Getenv("MYSQL_CONNECTION")
	DatabaseName          = "todo_list_app"
	JwtTimeOut            = time.Duration(6) * time.Hour
	JwtSigningMethod      = jwt.SigningMethodHS256
	JwtSignatureKey       = "test"
	// MongoConnectionString = os.Getenv("MONGO_CONNECTION")
)

const (
	AppName = "todo list app"
	BE_PORT = "8000"
)

func init() {
	if MysqlConnectionString == "" {
		MysqlConnectionString = "root:mysqlku123@tcp(127.0.0.1:3306)/todo_list_app?parseTime=true"
	}
}
