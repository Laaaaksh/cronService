package CRUD

import (
	"cronService/Database"
	"cronService/Models"
)

// getting all the cronjobs there in DB ---- perhaps we should add org_id to get cronjobs of certain org //

func GetCronJobs(CronJobs *[]Models.CronJob, username string) (err error) {

	var user Models.User

	if err := GetUserFromUserAuth(username, &user); err != nil {
		return err
	}
	if user.UserType == "Admin" {
		err = Database.DB.Where("organization_id = ?", user.OrganisationID).Find(CronJobs).Error
	} else {
		err = Database.DB.Where("user_id = ?", user.Id).Find(CronJobs).Error
	}
	if err != nil {
		return err
	}
	return nil
}

func GetCronJobById(job *Models.CronJob, id string) (err error) {
	if err = Database.DB.Where("id = ?", id).Find(job).Error; err != nil {
		return err
	}
	return nil
}
