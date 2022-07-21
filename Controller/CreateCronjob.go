package Controller

import (
	"cronService/Helpers"
	"cronService/Models"
	"cronService/Models/CRUD"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCronjob(c *gin.Context){

	jwttoken := c.Request.Header.Get("token")
	user_name,flag:=Helpers.ValidateToken(jwttoken)

	if !flag{
		c.JSON(http.StatusUnauthorized, gin.H{"error":"cannot access with provided token"})
		return
	}

	if ok := Helpers.CheckPermission(user_name, "create"); !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"You are not authorised to perform action"})
		return
	}

	var CronJob Models.CronJob
	err := c.ShouldBindJSON(&CronJob)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Please provide valid details"})
		return
	}

	if err = CRUD.CreateCronjob(&CronJob, user_name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"Problem creating cronjob"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message" : "success"})
	}
}