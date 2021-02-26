package model

import (
	mysqlConfig "GoPass/config/mysql"
	"GoPass/lib/helper"
	"GoPass/lib/mysql"
	"github.com/jinzhu/gorm"
)

type SystemMenu struct {
	Id          uint32          `json:"id"`
	Name        string          `gorm:"type:varchar(100);not null;comment('名称');" json:"name"`
	Path        string          `gorm:"type:varchar(50);not null;comment('路径');index" json:"path"`
	Component   string          `gorm:"type:varchar(100);not null;comment('组件');" json:"component"`
	Redirect    string          `gorm:"type:varchar(200);not null;comment('重定向');" json:"redirect"`
	Url         string          `gorm:"type:varchar(200);not null;comment('url');" json:"url"`
	MetaTitle   string          `gorm:"type:varchar(50);not null;comment('meta标题');" json:"meta_title"`
	MetaIcon    string          `gorm:"type:varchar(50);not null;comment('meta icon');" json:"meta_icon"`
	MetaNocache int             `gorm:"TINYINT(4);not null;default:0; comment('是否缓存（1:是- 0:否）') " json:"meta_nocache"`
	Alwaysshow  int             `gorm:"TINYINT(4);not null;default:0; comment('是否总是显示（1:是0：否）')" json:"alwaysshow"`
	MetaAffix   int             `gorm:"TINYINT(4);not null;default:0; comment('是否加固（1:是0：否）')" json:"meta_affix"`
	Type        int             `gorm:"TINYINT(4);not null;default:2;comment('类型(1:固定,2:权限配置,3特殊)') " json:"meta_affix"`
	Hidden      int             `gorm:"TINYINT(4);not null;default:0;comment('是否隐藏（0否1是）')" json:"hidden"`
	Pid         uint32          `gorm:"not null;default:0;index:idx_list;comment('父ID')" json:"pid"`
	Sort        int             `gorm:"not null;default:0;index:idx_list;comment('排序')" json:"sort"`
	Status      int8            `gorm:"TINYINT(4);not null;default:1;index:idx_list;comment('状态（0禁止1启动）')" json:"status"`
	Level       int8            `gorm:"TINYINT(4);not null;default:0;comment('层级')" json:"level"`
	Ctime       helper.JSONTime `gorm:"not null;default:current_timestamp" json:"ctime"`
}

func (m SystemMenu) Model() *gorm.DB { return mysql.Mysql((&m).Connection()).Model(&m) }

func (*SystemMenu) Connection() string { return mysqlConfig.Config.Default }

func (*SystemMenu) TableName() string { return "system_menu" }

func (m SystemMenu) TreeNode(menuMap map[uint32][]SystemMenu, pid uint32) []SystemMenu {
	var menuNewArr []SystemMenu
	if _, ok := menuMap[pid]; ok {
		for _, v := range menuMap[pid] {
			menuNewArr = append(menuNewArr, v)
			menuNewArr = append(menuNewArr, m.TreeNode(menuMap, v.Id)...)
		}
	}
	return menuNewArr
}

func (m SystemMenu) TreeMenuNew(menuMap map[uint32][]SystemMenu, pid uint32, mrArr map[uint32][]string) []interface{} {
	var menuNewArr []interface{}
	if _, ok := menuMap[pid]; ok {
		for _, value := range menuMap[pid] {
			var item = make(map[string]interface{})
			item["path"] = value.Path
			item["component"] = value.Component
			if value.Redirect != "" {
				item["redirect"] = value.Redirect
			}
			if value.Alwaysshow == 1 {
				item["alwaysShow"] = true
			}
			if value.Hidden == 1 {
				item["hidden"] = true
			} else {
				item["hidden"] = false
			}
			var meta = make(map[string]interface{})
			if _, ok := mrArr[value.Id]; ok {
				meta["roles"] = mrArr[value.Id]
			}
			if value.MetaTitle != "" {
				meta["title"] = value.MetaTitle
			}
			if value.MetaIcon != "" {
				meta["icon"] = value.MetaIcon
			}
			if value.MetaAffix == 1 {
				meta["affix"] = true
			}
			if value.MetaNocache == 1 {
				meta["noCache"] = true
			}
			if value.Status == 1 {
				meta["status"] = true
			}

			if len(meta) > 0 {
				item["meta"] = meta
			}
			item["pid"] = value.Pid
			item["id"] = value.Id
			item["url"] = value.Url
			item["name"] = value.Name
			children := m.TreeMenuNew(menuMap, value.Id, mrArr)
			if children != nil {
				item["children"] = children
			}
			menuNewArr = append(menuNewArr, item)
		}
	}
	return menuNewArr
}

func (m SystemMenu) GetRouteByRole(id interface{}) []SystemMenu {
	var constant []SystemMenu

	menu := SystemMenu{Type: 1}
	constant = menu.GetRowByType()

	var end []SystemMenu
	menu.Type = 3
	end = menu.GetRowByType()
	var async []SystemMenu
	async = menu.GetRowByRole(id)
	constant = append(constant, async...)
	constant = append(constant, end...)
	return constant
}

func (rm SystemMenu) GetRowByType() []SystemMenu {
	var systemmenus []SystemMenu
	rm.Model().Where("type=?", rm.Type).Find(&systemmenus)
	return systemmenus
}

func (m SystemMenu) GetRowByRole(id interface{}) []SystemMenu {
	var menu []SystemMenu
	m.Model().Select("system_menu.*").
		Joins(" join system_role_menu on system_role_menu.system_menu_id =  system_menu .id").
		Where("system_menu.status = ?", 1).
		Where("system_role_menu.system_role_id = ?", id).
		Find(&menu)
	return menu
}
