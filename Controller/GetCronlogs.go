package Controller

import (
	"cronService/Models"
	"cronService/Models/CRUD"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCronLogs(c gin.Context){
	var logs Models.CronJob

	Id := c.Params.ByName("id")
	err := CRUD.GetCronLogs(&logs,Id)
	if err != nil{
		c.JSON(http.StatusBadRequest, "Cron Job with the given id not found")
	}else {
		c.JSON(http.StatusOK, logs)
	}

}
