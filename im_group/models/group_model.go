package models

import (
	"im_server/common/models"
	"im_server/common/models/ctype"
)

// GroupModel
// @Description: 群组表
// @Author Oberl-Fitzgerald 2024-06-14 19:27:58
type GroupModel struct {
	models.Models
	Title              string                `gorm:"size:32" json:"title"`     // 群名称
	Abstract           string                `gorm:"size:128" json:"abstract"` // 群简介
	Avatar             string                `gorm:"size:256" json:"avatar"`   // 群头像
	CreatorID          uint                  `json:"creator_id"`               // 创建者ID 群主
	IsSearch           bool                  `json:"is_search"`                // 是否允许被搜索
	Verify             int8                  `json:"verify"`                   // 群验证：0-不允许任何人添加，1-允许任何人添加，2-需要验证消息，3-需要回答问题，4-需要正确回答问题
	VerifyQuestion     *ctype.VerifyQuestion `json:"verify_question"`          // 验证问题 为3、4时需要
	IsInvite           bool                  `json:"is_invite"`                // 是否允许被邀请
	IsTemporarySession bool                  `json:"is_temporary_session"`     // 是否开启临时会话
	IsProhibit         bool                  `json:"is_prohibit"`              // 是否全员禁言
	Size               int                   `json:"size"`                     // 群容量 20 100 200 1000 2000
}
