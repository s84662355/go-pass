package model

import (
    "GoPass/lib/helper"
    "GoPass/lib/mysql"
    "github.com/jinzhu/gorm"
)

type AppDevice struct {
    Id            uint64          `json:"id"`
    DeviceId      string          `gorm:"size:64;not null;default:'';unique_index;" json:"device_id"` //设备 id
    RegionId      int64           `gorm:"not null;default:0;" json:"region_id"`                       //单位 id
    RegionName    string          `gorm:"size:32;not null;default:'';" json:"region_name"`            //单位名称
    UserId        int64           `gorm:"not null;default:0;" json:"user_id"`
    PluginVersion string          `gorm:"size:32;not null;default:'';" json:"plugin_version"` //插件版本
    DeviceModel   string          `gorm:"size:32;not null;default:'';" json:"device_model"`   //手机型号
    SystemVersion string          `gorm:"size:32;not null;default:'';" json:"system_version"` //系统版本
    Imei          string          `gorm:"size:32;not null;default:'';" json:"imei"`
    MacAddress    string          `gorm:"size:32;not null;default:'';" json:"mac_address"` //mac 地址
    AndroidId     string          `gorm:"size:32;not null;default:'';" json:"android_id"`
    SerialId      string          `gorm:"size:32;not null;default:'';" json:"serial_id"` //序列号
    NetType       string          `gorm:"size:32;not null;default:'';" json:"net_type"`  //网络类型
    NetInfo       string          `gorm:"size:32;not null;default:'';" json:"net_info"`  //网络信息
    IsClone       bool            `gorm:"not null;default:0;" json:"is_clone"`           //是否多开
    ScreenWidth   int64           `gorm:"not null;default:0;" json:"screen_width"`
    ScreenHeight  int64           `gorm:"not null;default:0;" json:"screen_height"`
    DyVersion     string          `gorm:"size:32;not null;default:'';" json:"dy_version"`
    IsBusy        bool            `gorm:"not null;default:0;" json:"is_busy"` //1忙碌 0空闲
    OnLine        bool            `gorm:"not null;default:0;" json:"on_line"` //1在线 0不在线
    CreatedAt     helper.JSONTime `gorm:"not null;default:current_timestamp" json:"created_at"`
    UpdatedAt     helper.JSONTime `gorm:"not null;default:current_timestamp" json:"updated_at"`
}

func (m AppDevice) Model() *gorm.DB {
    return mysql.Mysql((&m).Connection()).Model(&m)
}

func (*AppDevice) Connection() string {
    return "default"
}

func (*AppDevice) TableName() string {
    return "media_app_device"
}
