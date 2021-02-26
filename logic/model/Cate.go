package model

import (
	mysqlConfig "GoPass/config/mysql"
	"GoPass/lib/helper"
	"GoPass/lib/mysql"
	"github.com/jinzhu/gorm"
)

type Cate struct {
	Id        uint32           `json:"id"`
	Name      string           `gorm:"type:varchar(100);not null;unique_index" json:"name"` //分类名称
	CreatedAt helper.JSONTime  `gorm:"not null;default:current_timestamp" json:"created_at"`
	UpdatedAt helper.JSONTime  `gorm:"Extra:on update CURRENT_TIMESTAMP   ;not null;default:current_timestamp;" json:"updated_at"`
	DeletedAt *helper.JSONTime `json:"deleted_at"`
}

func (m Cate) Model() *gorm.DB { return mysql.Mysql((&m).Connection()).Model(&m) }

func (*Cate) Connection() string { return mysqlConfig.Config.Default }

func (*Cate) TableName() string { return "cate" }

func (m Cate) All() []Cate {
	var res []Cate
	m.Model().Where("deleted_at is null").Find(&res)
	return res
}

func (m *Cate) Create() bool {
	m.CreatedAt = helper.JSONTime{}.Create()
	m.UpdatedAt = helper.JSONTime{}.Create()
	m.Model().Create(m)
	if m.Id == 0 {
		return false
	}
	return true
}

func (a *Cate) Update() {
	a.UpdatedAt = helper.JSONTime{}.Create()
	a.Model().Save(a)
}

func (a Cate) Get(id interface{}) Cate {
	res := Cate{}
	a.Model().Where("id = ?", id).First(&res)
	return res
}

func (a *Cate) Delete() {
	DeletedAt := helper.JSONTime{}.Create()
	a.DeletedAt = &DeletedAt
	a.Model().Save(a)
}
