package Controller

import (
	"cronService/Models/CRUD"
	"cronservice/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c * gin.Context){
	var user Models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Please provide valid details"})
		return
	}

	if err := CRUD.CreateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Please use a different User Name"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message" : "success"})
	}
}
