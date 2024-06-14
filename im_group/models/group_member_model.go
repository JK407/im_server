package models

import "im_server/common/models"

// GroupMemberModel
// @Description: 群成员表
// @Author Oberl-Fitzgerald 2024-06-14 19:44:45
type GroupMemberModel struct {
	models.Models
	GroupID        uint       `json:"group_id"`                        // 群ID
	GroupModel     GroupModel `gorm:"foreignKey:GroupID" json:"_"`     // 群外键
	UserID         uint       `json:"user_id"`                         // 用户ID
	MemberNickName string     `gorm:"size:32" json:"member_nick_name"` // 群内昵称
	Role           int8       `json:"role"`                            // 角色：1-群主，2-管理员，3-普通成员
	ProhibitTime   *int64     `json:"prohibit_time"`                   // 禁言时间 单位分钟
}
