package domain

type Role struct {
	ID       uint   `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"roleId"`
	RoleName string `gorm:"type:varchar(25) NOT NULL;" json:"roleName"`
}
