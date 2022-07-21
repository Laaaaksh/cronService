package Controller

import (
	"cronService/Models"
	"cronService/Models/CRUD"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCronJobs(c *gin.Context) {

	jwttoken := c.Request.Header.Get("token")
	user_name,flag:= CRUD.ValidateToken(jwttoken)

	if !flag{
		c.JSON(http.StatusUnauthorized, gin.H{"error":"cannot access with provided token"})
		return
	}

	var CronJobs []Models.CronJob

	err := CRUD.GetCronJobs(&CronJobs,user_name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"Problem getting cronjob"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message" : "success"})
	}
}

func GetCronJobById(c *gin.Context){
	jwttoken := c.Request.Header.Get("token")
	user_name,flag:= CRUD.ValidateToken(jwttoken)

	if !flag{
		c.JSON(http.StatusUnauthorized, gin.H{"error":"cannot access with provided token"})
		return
	}

	id := c.Params.ByName("id")

	if flag= CRUD.CheckPermission(user_name, "nil", id); !flag{
		c.JSON(http.StatusUnauthorized, gin.H{"error":"cannot access requested cronjob"})
		return
	}
	var job Models.CronJob
	if err:= CRUD.GetCronJobById(&job,id); err!= nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
	}
}