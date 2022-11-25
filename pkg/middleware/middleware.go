package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt/v4"
	"github.com/hasban-fardani/todo-list-app/pkg/auth"
	"github.com/hasban-fardani/todo-list-app/pkg/configs"
	"github.com/hasban-fardani/todo-list-app/pkg/models"
)

func NeedLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userData models.User

		// get token string
		tokenStr := ctx.GetHeader("Authorization")
		tokenStr = strings.ReplaceAll(tokenStr, "Bearer ", "")
		ok, err := auth.ValidateToken(tokenStr)
		if err != nil || !ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusInternalServerError,
				"token":  tokenStr,
				"error":  err.Error(),
			})
			ctx.Abort()
			return
		}

		// check token is valid
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("signing method invalid")
			} else if method != configs.JwtSigningMethod {
				return nil, fmt.Errorf("signing method invalid")
			}
			return []byte(configs.JwtSignatureKey), nil
		})
		if err != nil {
			fmt.Println(err.Error())
		}

		// decode jwt to user struct
		claims := token.Claims.(jwt.MapClaims)
		user := claims["user"].(map[string]interface{})
		userStr, err := json.Marshal(user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			ctx.Abort()
			return
		}
		json.Unmarshal(userStr, &userData)
		ctx.Set("user", userData)

		// next to route
		ctx.Next()
	}
}
