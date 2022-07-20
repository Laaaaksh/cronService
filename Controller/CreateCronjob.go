package Controller

import (
	"cronService/Models"
	"cronService/Models/CRUD"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCronjob(c *gin.Context){
	var CronJob Models.CronJob
	err := c.ShouldBindJSON(&CronJob)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Please provide valid details"})
		return
	}

	if err := CRUD.CreateCronjob(&CronJob); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"Problem creating cronjob"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message" : "success"})
	}
}