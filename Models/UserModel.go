package Models

import "time"

type User struct {
	Id             int       `json:"_id"`
	UserType       string    `json:"user_type"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	OrganisationId int       `json:"organisation_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAT      time.Time `json:"updated_at"`
}

func (u *User) TableName() string {
	return "user"
}
