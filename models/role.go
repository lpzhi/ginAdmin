package models


type Role struct {
	Model
	RoleName string  `json:"role_name"`
	creatOn int	`json:"creat_on"`
}

func GetRoleTotal() (count int)  {
	db.Model(&Role{}).Count(&count)
	return
}

func GetRoles(offset,pageSize int,maps interface{})  (roles []Role) {
	db.Debug().Where(maps).Offset(offset).Limit(pageSize).Find(&roles)
	return
}
