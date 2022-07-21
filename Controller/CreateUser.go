package Controller

import (
	"cronService/Helpers"
	"cronService/Models/CRUD"
	"github.com/gin-gonic/gin"
	"net/http"
)
type UserUserauth struct{
	UserType       string `json:"user_type"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	UserName     string `json:"user_name"`
	HashPassword string `json:"hash_password"`
}
func CreateUser(c * gin.Context){

	jwttoken := c.Request.Header.Get("token")
	user_name,flag:=Helpers.ValidateToken(jwttoken)

	if !flag{
		c.JSON(http.StatusUnauthorized, gin.H{"error":"cannot access with provided token"})
		return
	}
	if !CRUD.AuthorizeAdmin(user_name) {
		c.JSON(http.StatusForbidden, gin.H{"error":"unauthorized to create user"})
	}

	var user UserUserauth

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Please provide valid details"})
		return
	}
	id, err := Helpers.GetOrgID(user_name)
	if err := CRUD.CreateUser(&user, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Please use a different User Name"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message" : "success"})
	}
}
