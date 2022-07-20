package Models

type UserToken struct {
	Id                int    `json:"_id"`
	UserId            int    `json:"user_id"`
	Token             string `json:"token"`
	TokenExpiryTime   string `json:"token_expiry_time"`
	TokenCreationTime string `json:"token_creation_time"`
}

func (ut *UserToken) TableName() string {
	return "usertoken"
}
