package Controller

import (
	"cronService/Models"
	"cronService/Models/CRUD"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePermissions(c *gin.Context) {
	fmt.Println("ff")
	jwttoken := c.Request.Header.Get("token")
	user_name, flag := CRUD.ValidateToken(jwttoken)
	fmt.Println(user_name)
	if !flag {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "cannot access with provided token"})
		return
	}
	if !CRUD.AuthorizeAdmin(user_name) {
		c.JSON(http.StatusForbidden, gin.H{"error": "unauthorized to create Permission group"})
		return
	}
	var permission Models.PermissionType
	c.BindJSON(&permission)
	err := CRUD.CreatePermissions(&permission)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "failure"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}
