package model

import (
	mysqlConfig "GoPass/config/mysql"
	"GoPass/lib/mysql"
	"github.com/jinzhu/gorm"
)

type SystemRoleMenu struct {
	Id           uint32 `json:"id"`
	SystemRoleId uint32 `gorm:"not null;default:0;comment('角色主键');index:system_role_id" json:"system_role_id"`
	SystemMenuId uint32 `gorm:"not null;default:0;comment('菜单主键');index:system_role_id" json:"system_menu_id"`
}

func (m SystemRoleMenu) Model() *gorm.DB { return mysql.Mysql((&m).Connection()).Model(&m) }

func (*SystemRoleMenu) Connection() string { return mysqlConfig.Config.Default }

func (*SystemRoleMenu) TableName() string { return "system_role_menu" }
