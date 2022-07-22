package Controller

import (
	"cronService/Models"
	"cronService/Models/CRUD"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c * gin.Context){

	jwttoken := c.Request.Header.Get("token")
	user_name,flag:= CRUD.ValidateToken(jwttoken)
	fmt.Println(jwttoken)
	if !flag{
		c.JSON(http.StatusUnauthorized, gin.H{"error":"cannot access with provided token"})
		return
	}
	if !CRUD.AuthorizeAdmin(user_name) {
		c.JSON(http.StatusForbidden, gin.H{"error":"unauthorized to create user"})
		return
	}

	var user Models.UserUserauth

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Please provide valid details"})
		return
	}
	id, err := CRUD.GetOrgID(user_name)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := CRUD.CreateUser(&user, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message" : "success"})
	}
}
