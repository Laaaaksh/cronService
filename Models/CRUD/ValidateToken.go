package CRUD

import (
	"cronService/Models"
	"github.com/dgrijalva/jwt-go"
)

func ValidateToken(jwttoken string) (string, bool) {
	token, _ := jwt.ParseWithClaims(jwttoken, &Models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("sercrethatmaycontainch@r$32chars"), nil
	})
	claims, ok := token.Claims.(*Models.Claims)

	if ok && token.Valid {
		return claims.UserName, true
	}
	return "", false
}
