package Models

type PermissionUserType struct {
	Id               int    `json:"_id"`
	UserType         string `json:"user_type"`
	PermissionTypeId int    `json:"permission_type_id"`
}

func (p *PermissionUserType) TableName() string {
	return "permissionusertype"
}
