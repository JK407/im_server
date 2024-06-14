package models

import "im_server/common/models"

// UserConfModel
// @Description: 用户配置表
// @Author Oberl-Fitzgerald 2024-06-14 18:25:34
type UserConfModel struct {
	models.Models
	UserID         uint            `json:"user_id"`         // 用户ID
	RecallMessage  *string         `json:"recall_message"`  // 撤回消息
	FriendOnline   bool            `json:"friend_online"`   // 好友上线提醒
	Sound          bool            `json:"sound"`           // 声音提醒
	SecureLink     bool            `json:"secure_link"`     // 安全链接
	SavePwd        bool            `json:"save_pwd"`        // 保存密码
	Online         bool            `json:"online"`          // 在线状态
	SearchUser     int8            `json:"search_user"`     // 别人查找到你的方式：0-不允许别人查找，1-只通过用户号查找，2-可以通过昵称查找
	FriendVerify   int8            `json:"friend_verify"`   // 好友验证：0-不允许任何人添加，1-允许任何人添加，2-需要验证消息，3-需要回答问题，4-需要正确回答问题
	VerifyQuestion *VerifyQuestion `json:"verify_question"` // 验证问题 为3、4时需要
}

// VerifyQuestion
// @Description: 验证问题
// @Author Oberl-Fitzgerald 2024-06-14 18:25:50
type VerifyQuestion struct {
	Question1 *string `json:"question1"` // 问题1
	Question2 *string `json:"question2"` // 问题2
	Question3 *string `json:"question3"` // 问题3
	Answer1   *string `json:"answer1"`   // 答案1
	Answer2   *string `json:"answer2"`   // 答案2
	Answer3   *string `json:"answer3"`   // 答案3
}
