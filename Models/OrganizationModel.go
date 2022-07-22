package Models

type Organization struct {
	Id               int    `json:"id"`
	OrganisationName string `gorm:"unique" json:"organisation_name"`
	CreatedAt        int64  `json:"created_at"`
	UpdatedAT        int64  `json:"updated_at"`
}
type OrgUser struct {
	OrganisationName string `json:"organisation_name" binding:"required" gorm:"unique"`
	AdminUserName    string `json:"admin_user_name" binding:"required" gorm:"unique"`
	Password         string `json:"password" binding:"required"`
}

func (o *Organization) TableName() string {
	return "organization"
}
