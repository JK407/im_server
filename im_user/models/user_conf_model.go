package models

import (
	"im_server/common/models"
	"im_server/common/models/ctype"
)

// UserConfModel
// @Description: 用户配置表
// @Author Oberl-Fitzgerald 2024-06-14 18:25:34
type UserConfModel struct {
	models.Models
	UserID         uint                  `json:"user_id"`                       // 用户ID
	UserModel      UserModel             `gorm:"foreignKey:UserID" json:"_"`    // 外键
	RecallMessage  *string               `gorm:"size:32" json:"recall_message"` // 撤回消息
	FriendOnline   bool                  `json:"friend_online"`                 // 好友上线提醒
	Sound          bool                  `json:"sound"`                         // 声音提醒
	SecureLink     bool                  `json:"secure_link"`                   // 安全链接
	SavePwd        bool                  `json:"save_pwd"`                      // 保存密码
	Online         bool                  `json:"online"`                        // 在线状态
	SearchUser     int8                  `json:"search_user"`                   // 别人查找到你的方式：0-不允许别人查找，1-只通过用户号查找，2-可以通过昵称查找
	Verify         int8                  `json:"verify"`                        // 好友验证：0-不允许任何人添加，1-允许任何人添加，2-需要验证消息，3-需要回答问题，4-需要正确回答问题
	VerifyQuestion *ctype.VerifyQuestion `json:"verify_question"`               // 验证问题 为3、4时需要
}
