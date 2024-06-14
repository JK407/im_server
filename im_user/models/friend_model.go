package models

import "im_server/common/models"

// FriendModel
// @Description: 好友表
// @Author Oberl-Fitzgerald 2024-06-14 18:26:14
type FriendModel struct {
	models.Models
	SendUserID    uint      `json:"send_user_id"`                   // 发送验证方ID
	SendUserModel UserModel `gorm:"foreignKey:SendUserID" json:"_"` // 外键
	RevUserID     uint      `json:"rev_user_id"`                    // 接收验证方ID
	RevUserModel  UserModel `gorm:"foreignKey:RevUserID" json:"_"`  // 外键
	Notice        string    `gorm:"size:128" json:"notice"`         // 备注
}
