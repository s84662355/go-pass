package model

import (
    mysqlConfig "GoPass/config/mysql"
    "GoPass/lib/helper"
    "GoPass/lib/mysql"
    _ "fmt"
    "github.com/jinzhu/gorm"
)

type Setting struct {
    Id      uint32              `json:"id"`
    KeyName string              `gorm:"type:varchar(100);not null;unique_index" json:"key_name"`
    Content helper.JsonSqlValue `gorm:"type:text;not null;" json:"content"`
}

func (m Setting) Model() *gorm.DB {
    return mysql.Mysql((&m).Connection()).Model(&m)
}

func (*Setting) Connection() string {
    return mysqlConfig.Config.Default
}

func (*Setting) TableName() string {
    return "setting"
}
