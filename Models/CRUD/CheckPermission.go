package CRUD

import (
	"cronService/Database"
	"cronService/Models"
)

func CheckPermission(userName string, permission string, jobID string) bool {

	var user Models.User
	var job Models.CronJob
	if err := GetUserFromUserAuth(userName, &user); err != nil {
		return false
	}

	if jobID != "Create" {
		GetCronJobById(&job, jobID)

		if user.UserType == "Admin" {
			if user.OrganisationID == job.OrganizationId {
				return true
			} else {
				return false
			}
		}
		if user.Id != job.UserId {
			return false
		}
	}
	if user.UserType == "Admin" {
		return true
	}

	var permissionType Models.PermissionType
	if err := Database.DB.Where("id = ?", user.PermissionID).Find(&permissionType).Error; err != nil {
		return false
	}

	if permission == "delete" {
		if permissionType.Delete {
			return true
		} else {
			return false
		}
	} else if permission == "add" {
		if permissionType.Add {
			return true
		} else {
			return false
		}
	} else if permission == "update" {
		if permissionType.Update {
			return true
		} else {
			return false
		}
	} else if permission == "disable" {
		if permissionType.Disable {
			return true
		} else {
			return false
		}
	} else if permission == "enable" {
		if permissionType.Enable {
			return true
		} else {
			return false
		}
	} else if permission == "logs" {
		if permissionType.Logs {
			return true
		} else {
			return false
		}
	}
	return true

}
