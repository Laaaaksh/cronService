package Models

import "time"

type Organization struct {
	Id               int       `json:"_id"`
	OrganisationName string    `json:"organisation_name"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAT        time.Time `json:"updated_at"`
}

func (o *Organization) TableName() string {
	return "organization"
}
