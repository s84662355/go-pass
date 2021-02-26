package model

import (
	mysqlConfig "GoPass/config/mysql"
	"GoPass/lib/helper"
	"GoPass/lib/mysql"
	"github.com/jinzhu/gorm"
)

type SystemUser struct {
	Id            uint32          `json:"id"`
	Name          string          `gorm:"type:varchar(50);not null;comment('姓名')" json:"name"`
	Nickname      string          `gorm:"type:varchar(50);not null;comment('用户登录名');unique_index" json:"nickname"`
	Password      string          `gorm:"type:varchar(50);not null;comment('密码')" json:"password"`
	Salt          string          `gorm:"type:varchar(4);not null;comment('盐')" json:"salt"`
	Phone         string          `gorm:"type:varchar(11);not null;default:'';comment('手机号')" json:"phone"`
	Avatar        string          `gorm:"type:varchar(300);not null;default:'';comment('头像')" json:"avatar"`
	Introduction  string          `gorm:"type:varchar(300);not null;default:'';comment('简介')" json:"introduction"`
	Status        int8            `gorm:"TINYINT(4);not null;default:1;comment('状态（0 停止1启动）" json:"status"`
	Utime         helper.JSONTime `gorm:"TIMESTAMP;not null;default:current_timestamp;comment('更新时间')" json:"utime"`
	LastLoginTime helper.JSONTime `gorm:"DATETIME;not null;default:'0000-00-00 00:00:00';comment('上次登录时间')" json:"last_login_time"`
	LastLoginIp   string          `gorm:"type:varchar(50);not null;default:'';comment('最近登录IP')" json:"last_login_ip"`
	Ctime         helper.JSONTime `gorm:"not null;default:current_timestamp" json:"ctime"`
}

func (m SystemUser) Model() *gorm.DB { return mysql.Mysql((&m).Connection()).Model(&m) }

func (*SystemUser) Connection() string { return mysqlConfig.Config.Default }

func (*SystemUser) TableName() string { return "system_user" }

func (u SystemUser) GetAllByName(name string) []map[string]interface{} {
	var systemusers []map[string]interface{}
	u.Model().Where("name like ?", name+"%").Find(&systemusers)
	return systemusers
}

func (u *SystemUser) Add(roles []interface{}) (uint32, bool) {
	var tx *gorm.DB
	rollback := true
	tx = u.Model().Begin()

	defer func() {
		if rollback {
			tx.Rollback()
		}
	}()

	tx.Create(u)

	//如果没有设置权限
	if len(roles) < 1 {
		tx.Commit()
		rollback = false
		return u.Id, true
	}

	for _, k := range roles {
		roleModel := SystemRole{}

		tx.Where("name = ?", k.(string)).First(&roleModel)

		if roleModel.Id == 0 {
			continue
		}
		if roleModel.Status == 0 {
			continue
		}

		userroleModel := SystemUserRole{}

		tx.Where("system_role_id = ? and system_user_id = ?", roleModel.Id, u.Id).First(&userroleModel)

		if userroleModel.Id != 0 {
			continue
		}

		userroleModel = SystemUserRole{SystemRoleId: roleModel.Id, SystemUserId: u.Id}

		tx.Create(&userroleModel)

		if userroleModel.Id == 0 {
			return 0, false
		}
	}
	rollback = false
	tx.Commit()

	return u.Id, true
}

func (u *SystemUser) Update(roles []interface{}) bool {
	var tx *gorm.DB
	rollback := true
	tx = u.Model().Begin()

	defer func() {
		if rollback {
			tx.Rollback()
		}
	}()

	tx.Where("id = ?", u.Id).Update(u)

	tx.Where("system_user_id=?", u.Id).Delete(SystemUserRole{})

	//如果没有设置权限
	if len(roles) < 1 {
		rollback = false
		tx.Commit()
		return true
	}

	for _, k := range roles {
		roleModel := SystemRole{}

		tx.Where("name = ?", k.(string)).First(&roleModel)

		if roleModel.Id == 0 {
			continue
		}
		if roleModel.Status == 0 {
			continue
		}

		userroleModel := SystemUserRole{}

		tx.Where("system_role_id = ? and system_user_id = ?", roleModel.Id, u.Id).First(&userroleModel)

		if userroleModel.Id != 0 {
			continue
		}

		userroleModel = SystemUserRole{SystemRoleId: roleModel.Id, SystemUserId: u.Id}

		tx.Create(&userroleModel)

		if userroleModel.Id == 0 {
			return false
		}
	}
	rollback = false
	tx.Commit()
	return true
}
