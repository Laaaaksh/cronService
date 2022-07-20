package Routes

import (
	"cronService/Controller"
	"github.com/gin-gonic/gin"
)

// Routes

func Setuprouter() *gin.Engine {

	r := gin.Default()
	grp := r.Group("/cronjobs")
	{
		// Laksh Routes

		//dineshroutes

		//Harihar Routes

		grp.POST("change-status/:id", Controller.ChangeCronJobStatusByID)
		grp.POST("edit/:id", Controller.UpdateCronJobByID)
		grp.DELETE("delete/:id", Controller.DeleteCronJobByID)
	}
	return r

}


