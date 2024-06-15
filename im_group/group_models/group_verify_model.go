package group_models

import (
	"im_server/common/models"
	"im_server/common/models/ctype"
)

// GroupVerifyModel
// @Description: 群验证表
// @Author Oberl-Fitzgerald 2024-06-14 19:45:03
type GroupVerifyModel struct {
	models.Models
	GroupID           uint                  `json:"group_id"`                          // 群ID
	GroupModel        GroupModel            `gorm:"foreignKey:GroupID" json:"-"`       // 群外键
	UserID            uint                  `json:"user_id"`                           // 用户ID
	AdditionalMessage *string               `gorm:"size:32" json:"additional_message"` // 附加消息
	VerifyQuestion    *ctype.VerifyQuestion `json:"verify_question"`                   // 验证问题 为3、4时需要
	Status            int8                  `json:"status"`                            // 状态：0-未操作，1-同意，2-拒绝，3-忽略
	Type              int8                  `json:"type"`                              // 类型：1-申请加入群，2-退群
}
