package Models

type User struct {
	Id             int    `json:"id"`
	UserType       string `json:"user_type"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	OrganisationID int    `json:"organisation_id"`
	CreatedAt      int64  `json:"created_at"`
	UpdatedAT      int64  `json:"updated_at"`
	PermissionID   int    `json:"permission_id"`
}
type UserUserauth struct {
	UserType     string `json:"user_type"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	UserName     string `json:"user_name"`
	HashPassword string `json:"hash_password"`
}

func (u *User) TableName() string {
	return "user"
}
