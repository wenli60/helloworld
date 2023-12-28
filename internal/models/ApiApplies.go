package models

import (
	"gorm.io/gorm"
)

// ApiApplies xx
type ApiApplies struct {
	gorm.Model
	Appid             string `gorm:"type:varchar(16);comment:系统缩写;not null;" json:"app_id"`
	Appkey            string `gorm:"type:varchar(32);not null;comment:密钥" json:"app_key"`
	PlatformName      string `gorm:"type:varchar(255);not null;comment:对方平台名称" json:"platform_name"`
	PlatformUrl       string `gorm:"type:varchar(255);not null;comment:对方平台地址" json:"platform_url"`
	PlatformDeveloper string `gorm:"type:varchar(32); null;comment:对方平台地址" json:"platform_developer"`
	Status            int    `gorm:"type:int(1);not null;default:0;comment:状态" json:"status"`
	ApplyReason       string `gorm:"type:varchar(512);comment:申请理由" json:"apply_reason"`
}

func CheckAuth(appid, appkey string) bool {
	var apiApplies ApiApplies
	_ = GetDBConn().Where("`appid` = ? and appkey= ?", appid, appkey).Last(&apiApplies).Error

	if apiApplies.ID > 0 {
		return true
	}

	return false
}
