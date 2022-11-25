package auth

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/hasban-fardani/todo-list-app/pkg/configs"
	"github.com/hasban-fardani/todo-list-app/pkg/models"
)

func CreateUserToken(claims models.UserClaims) (string, error) {
	token := jwt.NewWithClaims(configs.JwtSigningMethod, claims)
	signedToken, err := token.SignedString([]byte(configs.JwtSignatureKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateToken(tokenStr string) (bool, error) {
	tokenStr = strings.ReplaceAll(tokenStr, "Bearer ", "")
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != configs.JwtSigningMethod {
			return nil, fmt.Errorf("signing method invalid")
		}
		return []byte(configs.JwtSignatureKey), nil
	})
	if err != nil {
		return false, err
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return false, fmt.Errorf("token not valid")
	}

	return true, nil
}
