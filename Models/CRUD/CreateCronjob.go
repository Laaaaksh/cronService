package CRUD

import (
	"cronService/Database"
	"cronService/Models"
)
// simply creating cj what all inputs do we get from user?
//and whats getting passed and autherization thing needs to be updated.

func CreateCronjob(CronJob *Models.CronJob) (err error)   {
	if err := Database.DB.Create(CronJob).Error; err != nil {
		return err
	}
	return nil
}
