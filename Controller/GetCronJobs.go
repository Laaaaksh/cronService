package Controller

import (
	"cronService/Helpers"
	"cronService/Models"
	"cronService/Models/CRUD"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCronJobs(c *gin.Context) {

	jwttoken := c.Request.Header.Get("token")
	user_name,flag:=Helpers.ValidateToken(jwttoken)

	if !flag{
		c.JSON(http.StatusUnauthorized, gin.H{"error":"cannot access with provided token"})
		return
	}

	var CronJobs []Models.CronJob

	err := CRUD.GetCronJobs(&CronJobs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"Problem getting cronjob"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message" : "success"})
	}
}