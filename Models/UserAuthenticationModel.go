package Models

type UserAuthentication struct {
	Id           int    `json:"_id"`
	UserName     string `json:"user_name"`
	HashPassword string `json:"hash_password"`
	UserID       int    `json:"user_id"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAT    int64  `json:"updated_at"`
}

func (ua *UserAuthentication) TableName() string {
	return "userauthentication"
}
