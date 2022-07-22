package Models

import "github.com/dgrijalva/jwt-go"

type UserAuthentication struct {
	Id           int    `json:"_id"`
	UserName     string `json:"user_name" gorm:"unique"`
	HashPassword string `json:"hash_password"`
	UserID       int    `json:"user_id"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAT    int64  `json:"updated_at"`
}
type Credentials struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Claims struct {
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

func (ua *UserAuthentication) TableName() string {
	return "userauthentication"
}
