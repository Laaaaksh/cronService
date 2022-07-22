package Models

type PermissionType struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Add     bool   `json:"add"`
	Delete  bool   `json:"delete"`
	Update  bool   `json:"update"`
	Enable  bool   `json:"enable"`
	Disable bool   `json:"disable"`
	Logs    bool   `json:"logs"`
}

func (pt *PermissionType) TableName() string {
	return "permissiontype"
}
