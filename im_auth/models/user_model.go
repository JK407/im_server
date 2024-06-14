package models

import "im_server/common/models"

// UserModel
// @Description: 用户表
// @Author Oberl-Fitzgerald 2024-06-14 18:25:59
type UserModel struct {
	models.Models
	Pwd      string `gorm:"size:64" json:"pwd"`       // 密码
	NickName string `gorm:"size:32" json:"nick_name"` // 昵称
	Avatar   string `gorm:"size:256" json:"avatar"`   // 头像
	Abstract string `gorm:"size:128" json:"abstract"` // 简介
	IP       string `gorm:"size:32" json:"ip"`        // IP
	Addr     string `gorm:"size:64" json:"addr"`      // 地址
}
