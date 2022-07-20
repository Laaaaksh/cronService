package Middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/golang_heroku/common/response"
	"github.com/ydhnwb/golang_heroku/service"
	"log"
	"net/http"
)

func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc{
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == ""{
			res:=response.BuildErrorResponse("Failed to process the request", "No token provided", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest,res)
			return
		}

		token := jwtService.ValidateToken(authHeader,c)
		if token.Valid {
			claim := token.Claims.(jwt.MapClaims)
			// do something with claim
		} else {
			response := response.BuildErrorResponse("Error", "Invalid token", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
