package models

import "im_server/common/models"

// FriendModel
// @Description: 好友表
// @Author Oberl-Fitzgerald 2024-06-14 18:26:14
type FriendModel struct {
	models.Models
	SendUserID uint    `json:"send_user_id"` // 发送验证方ID
	RevUserID  uint    `json:"rev_user_id"`  // 接收验证方ID
	Notice     *string `json:"notice"`       // 备注

}
