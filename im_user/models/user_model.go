package models

import "im_server/common/models"

// UserModel
// @Description: 用户表
// @Author Oberl-Fitzgerald 2024-06-14 18:25:59
type UserModel struct {
	models.Models
	Pwd      string `json:"pwd"`       // 密码
	NickName string `json:"nick_name"` // 昵称
	Avatar   string `json:"avatar"`    // 头像
	Abstract string `json:"abstract"`  // 简介
	IP       string `json:"ip"`        // IP
	Addr     string `json:"addr"`      // 地址
}
