package main

import (
  "os"
  "os/signal"
  "syscall"

  "GoPass/config"
  "GoPass/lib/mysql"
  "GoPass/logic/model"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  //  "time"
)

func main() {
  done := make(chan os.Signal, 1)
  signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
  //连接mysql
  conn := mysql.ConnectMysql(config.MySQL, "default")
  defer mysql.DisconnectMysql()

  //迁移数据表
  //conn.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='终端设备表，保存每一个注册了的设备' ").AutoMigrate(&model.AppDevice{})

  // conn.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='设备与帐号绑定的索引表'").AutoMigrate(&model.DeviceAccountIndex{})

  conn.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='文章表' ").AutoMigrate(&model.Article{})
  conn.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='文章分类表' ").AutoMigrate(&model.Cate{})
  conn.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='配置表' ").AutoMigrate(&model.Setting{})

  /// <-done
}
