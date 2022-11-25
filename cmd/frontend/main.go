package frontend

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode("release")

	PORT := os.Getenv("FrontendPort")

	server := gin.New()
	server.GET("/", Index())

	server.Run(":" + PORT)
}
