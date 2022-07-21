package Controller

import (
	"cronService/Models"
	"cronService/Models/CRUD"
	"github.com/gin-gonic/gin"
	"net/http"
)


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
