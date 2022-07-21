package CRUD

import (
	"cronService/Models"
	"github.com/dgrijalva/jwt-go"
)

func ValidateToken(jwttoken string) (string, bool){
	token, _ := jwt.ParseWithClaims(jwttoken, &Models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"), nil
	})
	claims, ok := token.Claims.(*Models.Claims)

	if !ok || !token.Valid {
		return "",false
	}
	return claims.UserName, true
}
