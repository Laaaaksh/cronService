package Controller
import (
	"cronService/Models"
	"cronService/Models/CRUD"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserPermissionTypes(c *gin.Context){

	var permission Models.PermissionType
	c.BindJSON(&permission)
	err:= CRUD.UserPermissionTypes(&permission)
	if err!=nil{
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest,gin.H{"message":"failure"})
	}else{
		c.JSON(http.StatusOK,gin.H{"message":"failure"})
	}

}
