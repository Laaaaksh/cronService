package Routes

import (
	"cronService/Controller"
	"github.com/gin-gonic/gin"
)

// Routes

func Setuprouter() *gin.Engine {

	r := gin.Default()
	grp1 := r.Group("/cronjobs")
	{
		grp1.POST("/create-cronjob/", Controller.CreateCronjob)
		grp1.POST("/change-status/:id", Controller.ChangeCronJobStatusByID)
		grp1.POST("/edit/:id", Controller.UpdateCronJobByID)
		grp1.DELETE("/delete/:id", Controller.DeleteCronJobByID)
		grp1.GET("/get-all-cronjobs", Controller.GetCronJobs)
		grp1.GET("/get-cronjob/:id", Controller.GetCronLogsById)
	}
	grp2 := r.Group("/user")
	{
		grp2.POST("/create-user/", Controller.CreateUser)
		grp2.POST("/login/", Controller.UserLogin)
		grp2.DELETE("/delete/:id", Controller.DeleteUserById)
		grp2.POST("/update-user/",Controller.UpdateUser)
	}

	grp3 := r.Group("/admin")
	{
		grp3.POST("/create-org/", Controller.CreateOrganization)
		grp3.POST("/create-permission/", Controller.CreatePermissions)
		grp3.GET("/get-permissions/", Controller.GetPermissions)

	}

	return r

}
