package model

import (
	mysqlConfig "GoPass/config/mysql"
	"GoPass/lib/helper"
	"GoPass/lib/mysql"
	"github.com/jinzhu/gorm"
)

type SystemLog struct {
	Id            uint64          `json:"id"`
	SystemUserId  uint32          `gorm:"not null;default:0;" json:"system_user_id"`
	Title         string          `gorm:"type:varchar(300);not null;" json:"title"`
	Content       string          `gorm:"type:text;not null;" json:"content"`
	RelationId    uint32          `gorm:"not null;default:0;index;" json:"relation_id"`
	RelationTable int             `gorm:"not null;default:1;index;comment('对应表(1 system_user,2 system_menu,3 system_role)');" json:"relation_table"`
	Ip            string          `gorm:"type:varchar(50);not null;" json:"ip"`
	Url           string          `gorm:"type:varchar(500);not null;" json:"url"`
	Ctime         helper.JSONTime `gorm:"not null;default:current_timestamp" json:"ctime"`
}

func (m SystemLog) Model() *gorm.DB { return mysql.Mysql((&m).Connection()).Model(&m) }

func (*SystemLog) Connection() string { return mysqlConfig.Config.Default }

func (*SystemLog) TableName() string { return "system_log" }
