package Controller

import (
	"cronService/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ChangeCronJobStatusByID(c *gin.Context){
      var cronjob Models.CronJob
      id := c.Params.ByName("id")

	  err := Models.GetCronJobByID(&cronjob, id)   // checking if there is any cronjob by this id
	  if err != nil {                              // if cronjob is not presented abort with failed status
	  	c.JSON(http.StatusNotFound, cronjob)
	  }

	  c.BindJSON(&cronjob)                         // else continue with changing the status
	  err := Models.ChangeCronJobStatus(&cronjob)
	  if err != nil {
	  	c.AbortWithStatus(http.StatusNotFound)
	  } else{
	  	c.JSON(http.StatusOK, cronjob)
	  }


}

func UpdateCronJobByID(c *gin.Context){

	var cronjob Models.CronJob
	id := c.Params.ByName("id")

	err := Models.GetCronJobByID(&cronjob, id)   // checking if there is any cronjob by this id
	if err != nil {                              // if cronjob is not presented abort with failed status
		c.JSON(http.StatusNotFound, cronjob)
	}

	c.BindJSON(&cronjob)                         // else continue with changing the status
	err := Models.UpdateCronJob(&cronjob)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else{
		c.JSON(http.StatusOK, cronjob)
	}



}

func DeleteCronJobByID(c * gin.Context){
		var cronjob Models.CronJob
		id := c.Params.ByName("id")
		err := Models.DeleteCronJob(&cronjob, id)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
		}
}

