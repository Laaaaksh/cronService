package Controller

import (
	"cronService/Models"
	"cronService/Models/CRUD"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ChangeCronJobStatusByID(c *gin.Context) {
	jwttoken := c.Request.Header.Get("token")
	user_name, flag := CRUD.ValidateToken(jwttoken)

	if !flag {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "cannot access with provided token"})
		return
	}
	type status struct {
		status string
	}
	var enable status
	var flag2 bool
	c.BindJSON(&enable)
	id := c.Params.ByName("id")
	if enable.status == "enable" {
		flag2 = CRUD.CheckPermission(user_name, "enable", id)
	} else {
		flag2 = CRUD.CheckPermission(user_name, "disable", id)
	}
	if !flag2 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to enable/disbale a cronjob"})
		return
	}
	var cronjob Models.CronJob

	err := CRUD.GetCronJobById(&cronjob, id) // checking if there is any cronjob by this id
	if err != nil {                          // if cronjob is not presented abort with failed status
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	c.BindJSON(&cronjob) // else continue with changing the status
	err = CRUD.ChangeCronJobStatus(&cronjob)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal server error"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}

}

func UpdateCronJobByID(c *gin.Context) {
	jwttoken := c.Request.Header.Get("token")
	user_name, flag := CRUD.ValidateToken(jwttoken)

	if !flag {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "cannot access with provided token"})
		return
	}
	id := c.Params.ByName("id")
	if flag = CRUD.CheckPermission(user_name, "update", id); !flag {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to update a cronjob"})
		return
	}

	var cronjob Models.CronJob

	err := CRUD.GetCronJobById(&cronjob, id) // checking if there is any cronjob by this id
	if err != nil {                          // if cronjob is not presented abort with failed status
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
	}

	c.BindJSON(&cronjob) // else continue with changing the status
	err = CRUD.UpdateCronJob(&cronjob)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal server error"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}

}

func DeleteCronJobByID(c *gin.Context) {

	jwttoken := c.Request.Header.Get("token")
	user_name, flag := CRUD.ValidateToken(jwttoken)

	if !flag {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "cannot access with provided token"})
		return
	}
	id := c.Params.ByName("id")
	if flag = CRUD.CheckPermission(user_name, "delete", id); !flag {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to delete a cronjob"})
		return
	}

	var cronjob Models.CronJob
	err := CRUD.DeleteCronJob(&cronjob, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal server error"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
