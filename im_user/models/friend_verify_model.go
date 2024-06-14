package models

import "im_server/common/models"

// FriendVerifyModel
// @Description: 好友验证表
// @Author Oberl-Fitzgerald 2024-06-14 18:33:41
type FriendVerifyModel struct {
	models.Models
	SendUserID        uint            `json:"send_user_id"`       // 发送验证方ID
	RevUserID         uint            `json:"rev_user_id"`        // 接收验证方ID
	Status            int8            `json:"status"`             // 状态：0-未操作，1-同意，2-拒绝，3-忽略
	AdditionalMessage *string         `json:"additional_message"` // 附加消息
	VerifyQuestion    *VerifyQuestion `json:"verify_question"`    // 验证问题 为3、4时需要
}
