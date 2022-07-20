package Models

import "time"

type UserAuthentication struct {
	Id           int       `json:"_id"`
	UserName     string    `json:"user_name"`
	HashPassword string    `json:"hash_password"`
	UserID       int       `json:"user_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAT    time.Time `json:"updated_at"`
}

func (ua *UserAuthentication) TableName() string {
	return "userauthentication"
}
