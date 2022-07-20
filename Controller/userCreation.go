package Controller

import (
	"cronService/Models"
	"cronService/Models/CRUD"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var jwtKey = []byte("my_secret_key")

type Credentials struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Claims struct {
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

func UserLogin(c *gin.Context) {
	var creds Credentials
	if err := c.ShouldBindJSON(&creds); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	if !CRUD.VerifyCredentials(creds) {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"invalid credentials"})
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		UserName: creds.UserName,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token":tokenString})

	var  userToken Models.UserToken
	userToken.UserName = creds.UserName
	userToken.TokenCreationTime = time.Now().Unix()
	userToken.TokenExpiryTime = expirationTime.Unix()
	userToken.Token = tokenString

	if err:= CRUD.CreateToken(&userToken); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err})
		return
	}
}


func CreateUserTypes(c *gin.Context) {
	var UserType Models.PermissionUserType

	if err := c.ShouldBindJSON(&UserType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := CRUD.CreateUserType(&UserType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
