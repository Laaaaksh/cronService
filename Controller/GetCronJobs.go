package Controller

import (
	"cronService/Models"
	"cronService/Models/CRUD"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCronJobs(c *gin.Context) {
	var CronJobs []Models.CronJob

	err := CRUD.GetCronJobs(&CronJobs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"Problem getting cronjob"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message" : "success"})
	}
}