package Models

type Organization struct {
	Id               int    `json:"_id"`
	OrganisationName string `json:"organisation_name"`
	CreatedAt        int64  `json:"created_at"`
	UpdatedAT        int64  `json:"updated_at"`
}

func (o *Organization) TableName() string {
	return "organization"
}
