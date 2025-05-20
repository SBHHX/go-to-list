package db

import (
	"go_to_list/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 修改Connect函数，接收dsn参数
func Connect(dsn string) error {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	// 自动迁移表结构
	return DB.AutoMigrate(&models.User{}, &models.List{}, &models.Task{})
}
