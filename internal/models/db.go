package models

import (
	tgorm "git.code.oa.com/trpc-go/trpc-database/gorm"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

// Setup 初始化数据库连接
func Setup() error {
	var err error
	dbConn, err = tgorm.NewClientProxy("trpc.mysql.tyoffice.notification")
	if err != nil {
		return err
	}
	return nil
}

// GetDBConn 获取主数据库的数据库连接
func GetDBConn() *gorm.DB {
	return dbConn
}
