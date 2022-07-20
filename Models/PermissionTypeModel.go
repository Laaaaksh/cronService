package Models

type PermissionType struct {
	Id      int    `json:"_id"`
	Name    string `json:"name"`
	Add     bool   `json:"add"`
	Delete  bool   `json:"delete"`
	Update  bool   `json:"update"`
	Enable  bool   `json:"enable"`
	Disable bool   `json:"disable"`
}

func (pt *PermissionType) TableName() string {
	return "permissiontype"
}
