package model

import (
	mysqlConfig "GoPass/config/mysql"
	"GoPass/lib/helper"
	"GoPass/lib/mysql"
	"github.com/jinzhu/gorm"
	"strings"
)

type SystemRole struct {
	Id          uint32          `json:"id"`
	Name        string          `gorm:"type:varchar(100);not null;comment('角色名称')" json:"name"`
	AliasName   string          `gorm:"type:varchar(50);not null;comment('别名')" json:"alias_name"`
	Description string          `gorm:"type:varchar(200);not null;comment('描述')" json:"description"`
	Status      int8            `gorm:"TINYINT(4);index;not null;default:1;comment('角色状态（0无效1有效）')" json:"status"`
	Type        int8            `gorm:"TINYINT(4);index;not null;default:1;comment('属于哪个应用')" json:"type"`
	Ctime       helper.JSONTime `gorm:"not null;default:current_timestamp" json:"ctime"`
}

func (m SystemRole) Model() *gorm.DB { return mysql.Mysql((&m).Connection()).Model(&m) }

func (*SystemRole) Connection() string { return mysqlConfig.Config.Default }

func (*SystemRole) TableName() string { return "system_role" }

func (r *SystemRole) GetRowMenu() map[uint32][]string {
	var sr []SystemRole
	r.Model().Find(&sr)

	var srMap map[uint32]string
	srMap = make(map[uint32]string, 0)
	for _, v := range sr {
		srMap[v.Id] = v.Name
	}
	var srm = SystemRoleMenu{}
	var rmArr []SystemRoleMenu
	srm.Model().Find(&rmArr)
	var mrMap = make(map[uint32][]string, 0)
	for _, value := range rmArr {
		_, ok := srMap[value.SystemRoleId]
		if ok {
			mrMap[value.SystemMenuId] = append(mrMap[value.SystemMenuId], srMap[value.SystemRoleId])
		}
	}
	return mrMap
}

func (r SystemRole) TreeRoutes(routes []interface{}) []uint32 {
	var ids []uint32
	for _, value := range routes {
		ids = append(ids, uint32(value.(map[string]interface{})["id"].(float64)))
		if _, ok := value.(map[string]interface{})["children"]; ok {
			children := value.(map[string]interface{})["children"].([]interface{})
			ids = append(ids, r.TreeRoutes(children)...)
		}
	}
	return ids
}

func (r *SystemRole) Update(data []uint32) bool {

	var tx *gorm.DB
	rollback := true
	tx = r.Model().Begin()
	defer func() {
		if rollback {
			tx.Rollback()
		}
	}()
	tx.Model(r).Updates(*r)
	if len(data) <= 0 {
		rollback = false
		tx.Commit()
		return true
	}

	rolemenu := SystemRoleMenu{SystemRoleId: r.Id}

	tx.Delete(&rolemenu)

	for _, value := range data {
		rm := SystemRoleMenu{SystemRoleId: r.Id, SystemMenuId: value}
		tx.Create(&rm)
		if rm.Id == 0 {
			return false
		}
	}
	rollback = false
	tx.Commit()
	return true
}

func (r *SystemRole) AddCommit(data []interface{}) bool {

	var tx *gorm.DB
	rollback := true
	tx = r.Model().Begin()

	defer func() {
		if rollback {
			tx.Rollback()
		}
	}()

	tx.Create(r)

	if r.Id == 0 {
		return false
	}
	if len(data) <= 0 {
		rollback = false
		tx.Commit()
		return true
	}
	for _, value := range data {

		menu := SystemMenu{}
		pathMain := value.(map[string]interface{})["path"].(string)
		menu.Path = pathMain
		menu.Component = value.(map[string]interface{})["component"].(string)
		menu.Type = 2
		initMenu := SystemMenu{}

		tx.Where("path=?", menu.Path).Where("component=?", menu.Component).Where("type=?", menu.Type).First(&initMenu)
		if initMenu.Id == 0 {
			continue
		}

		rm := SystemRoleMenu{SystemRoleId: r.Id, SystemMenuId: initMenu.Id}
		tx.Create(&rm)

		if rm.Id == 0 {
			return false
		}

		children := value.(map[string]interface{})["children"]
		if children == nil {
			continue
		}
		for _, v := range children.([]interface{}) {
			menu := SystemMenu{}
			strings.TrimPrefix(v.(map[string]interface{})["path"].(string), pathMain+"/")
			menu.Component = v.(map[string]interface{})["component"].(string)
			menu.Type = 2
			initMenu := SystemMenu{}

			tx.Where("path=?", menu.Path).Where("component=?", menu.Component).Where("type=?", menu.Type).First(&initMenu)
			if initMenu.Id == 0 {
				continue
			}
			rm := SystemRoleMenu{SystemRoleId: r.Id, SystemMenuId: initMenu.Id}
			tx.Create(&rm)

			if rm.Id == 0 {
				return false
			}
		}
	}
	rollback = false
	tx.Commit()
	return true
}

func (r *SystemRole) Delete() bool {
	var tx *gorm.DB
	rollback := true
	tx = r.Model().Begin()

	defer func() {
		if rollback {
			tx.Rollback()
		}
	}()

	r.Status = 0
	tx.Save(r)
	rollback = false
	tx.Commit()
	return true

}

func (r *SystemRole) GetNameList() []string {
	var list []string

	r.Model().Where("status = 1").Pluck("name", &list)

	return list
}
