package Controller

import (
	"cronService/Models/CRUD"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrgUser struct {
	OrganisationName string `json:"organisation_name"`
	AdminUserName    string `json:"admin_user_name"`
	Password         string `json:"password"`
}

func CreateOrganization(c *gin.Context) {
	var orguser OrgUser
	err := c.BindJSON(&orguser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Please provide valid details"})
		return
	}
	CRUD.CreateOrganization(orguser)
}
