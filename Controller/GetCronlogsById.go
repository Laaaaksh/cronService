package Controller

import (
	"cronService/Models"
	"cronService/Models/CRUD"
	"github.com/gin-gonic/gin"
	"net/http"
)


func GetCronLogsById(c *gin.Context){

	jwttoken := c.Request.Header.Get("token")
	user_name,flag:= CRUD.ValidateToken(jwttoken)

	if !flag{
		c.JSON(http.StatusUnauthorized, gin.H{"error":"cannot access with provided token"})
		return
	}
	Id := c.Params.ByName("id")
	flag2:= CRUD.CheckPermission(user_name,"logs",Id)
	if !flag2{
		c.JSON(http.StatusUnauthorized, gin.H{"error":"cannot access cron logs"})
		return
	}
	var logs Models.CronExecutionResult

	err := CRUD.GetCronLogsById(&logs,Id)
	if err != nil{
		c.JSON(http.StatusBadRequest, "Cron Job with the given id not found")
    return
	}else {
		c.JSON(http.StatusOK, logs)
    return
	}

}
