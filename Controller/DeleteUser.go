package Controller

import (
	"cronService/Models"
	"cronService/Models/CRUD"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteUserById(c *gin.Context) {
	jwttoken := c.Request.Header.Get("token")
	user_name, flag := CRUD.ValidateToken(jwttoken)

	if !flag {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "cannot access with provided token"})
		return
	}
	if !CRUD.AuthorizeAdmin(user_name) {
		c.JSON(http.StatusForbidden, gin.H{"error": "unauthorized to Delete user"})
	}

	var user Models.User

	id := c.Params.ByName("id")
	err := CRUD.DeleteUserById(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
