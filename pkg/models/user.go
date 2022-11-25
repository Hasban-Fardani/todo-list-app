package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	Id           int       `json:"id"`
	NamaLengkap  string    `json:"nama_lengkap"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	CreateAt     time.Time `json:"create_at"`
	LastEditedAt time.Time `json:"last_edited_at"`
}

type UserClaims struct {
	User `json:"user"`
	jwt.StandardClaims
}
