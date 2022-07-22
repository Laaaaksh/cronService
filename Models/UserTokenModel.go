package Models

type UserToken struct {
	Id                int    `json:"id"`
	UserName          string `json:"user_name"`
	Token             string `json:"token"`
	TokenExpiryTime   int64  `json:"token_expiry_time"`
	TokenCreationTime int64  `json:"token_creation_time"`
}

func (ut *UserToken) TableName() string {
	return "usertoken"
}
