package dao

import (
	"fmt"
	"go-web-play/setting"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 👉🏻 https://gorm.io/zh_CN/docs/

var (
	DB *gorm.DB
)

func InitMySQL(cfg *setting.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func Close() {
	err := DB.Close()
	if err != nil {
		return
	}
}
