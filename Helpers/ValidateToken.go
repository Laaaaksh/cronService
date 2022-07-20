package Helpers

import (
	"cronService/Controller"
	"github.com/dgrijalva/jwt-go"
)

func ValidateToken(jwttoken string) (string, bool){
	token, _ := jwt.ParseWithClaims(jwttoken, &Controller.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"), nil
	})
	claims, ok := token.Claims.(*Controller.Claims)

	if !ok || !token.Valid {
		return "",false
	}
	return claims.UserName, true
}
