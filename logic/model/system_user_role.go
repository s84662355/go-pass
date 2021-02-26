package model

import (
	mysqlConfig "GoPass/config/mysql"
	"GoPass/lib/mysql"
	"github.com/jinzhu/gorm"
)

type SystemUserRole struct {
	Id           uint32 `json:"id"`
	SystemUserId uint32 `gorm:"not null;default:0;index:system_user_id;comment('用户主键')" json:"system_user_id"`
	SystemRoleId uint32 `gorm:"not null;default:0;index:system_user_id;comment('角色主键')" json:"system_role_id"`
}

func (m SystemUserRole) Model() *gorm.DB   { return mysql.Mysql((&m).Connection()).Model(&m) }
func (*SystemUserRole) Connection() string { return mysqlConfig.Config.Default }
func (*SystemUserRole) TableName() string  { return "system_user_role" }
