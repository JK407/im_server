package models

import (
	"im_server/common/models"
	"time"
)

// ChatModel
// @Description: 聊天表
// @Author Oberl-Fitzgerald 2024-06-14 18:39:58
type ChatModel struct {
	models.Models
	SenderUserID uint       `json:"sender_user_id"` // 发送方ID
	RevUserID    uint       `json:"rev_user_id"`    // 接收方ID
	MsgType      int8       `json:"msg_type"`       // 消息类型：1-文本，2-图片，3-视频，4-文件，5-语音， 6-语音通话， 7-视频通话， 8-撤回消息， 9-回复消息， 10-引用消息， 11-转发消息
	MsgPreview   string     `json:"msg_preview"`    // 消息预览
	MsgContent   MsgContent `json:"msg_content"`    // 消息内容
	SystemMsg    *SystemMsg `json:"system_msg"`     // 系统提示
}

// SystemMsg
// @Description: 系统提示
// @Author Oberl-Fitzgerald 2024-06-14 19:06:09
type SystemMsg struct {
	Type int8 `json:"type"` // 违规类型 1-色情，2-广告，3-政治，4-诈骗，5-其他
}

// MsgContent
// @Description: 消息内容
// @Author Oberl-Fitzgerald 2024-06-14 18:53:41
type MsgContent struct {
	Type         int8                 `json:"type"`           // 消息类型：1-文本，2-图片，3-视频，4-文件，5-语音， 6-语音通话， 7-视频通话， 8-撤回消息， 9-回复消息， 10-引用消息， 11-转发消息
	Content      *string              `json:"content"`        // 文本内容 为1时使用
	ImageMsg     *ImageMsgContent     `json:"image_msg"`      // 图片内容 为2时使用
	VideoMsg     *VideoMsgContent     `json:"video_msg"`      // 视频内容 为3时使用
	FileMsg      *FileMsgContent      `json:"file_msg"`       // 文件内容 为4时使用
	VoiceMsg     *VoiceMsgContent     `json:"voice_msg"`      // 语音内容 为5时使用
	VoiceCallMsg *VoiceCallMsgContent `json:"voice_call_msg"` // 语音通话内容 为6时使用
	VideoCallMsg *VideoCallMsgContent `json:"video_call_msg"` // 视频通话内容 为7时使用
	WithdrawMsg  *WithdrawMsgContent  `json:"withdraw_msg"`   // 撤回消息内容 为8时使用
	ReplyMsg     *ReplyMsgContent     `json:"reply_msg"`      // 回复消息内容 为9时使用
	QuoteMsg     *QuoteMsgContent     `json:"quote_msg"`      // 引用消息内容 为10时使用
	ForwardMsg   *ForwardMsgContent   `json:"forward_msg"`    // 转发消息内容 为11时使用
}

// ImageMsgContent
// @Description: 图片消息内容
// @Author Oberl-Fitzgerald 2024-06-14 18:53:45
type ImageMsgContent struct {
	Title string `json:"title"` // 标题
	URL   string `json:"url"`   // 地址
}

// VideoMsgContent
// @Description: 视频消息内容
// @Author Oberl-Fitzgerald 2024-06-14 18:53:47
type VideoMsgContent struct {
	Title string `json:"title"` // 标题
	URL   string `json:"url"`   // 地址
	Time  int    `json:"time"`  // 时长 单位秒
}

// FileMsgContent
// @Description: 文件消息内容
// @Author Oberl-Fitzgerald 2024-06-14 18:53:51
type FileMsgContent struct {
	Title string `json:"title"` // 标题
	URL   string `json:"url"`   // 地址
	Size  int64  `json:"size"`  // 大小 单位字节
	Type  string `json:"type"`  // 文件类型
}

// VoiceMsgContent
// @Description: 语音消息内容
// @Author Oberl-Fitzgerald 2024-06-14 18:53:53
type VoiceMsgContent struct {
	URL  string `json:"url"`  // 地址
	Time int    `json:"time"` // 时长 单位秒
}

// VoiceCallMsgContent
// @Description: 语音通话消息内容
// @Author Oberl-Fitzgerald 2024-06-14 18:54:12
type VoiceCallMsgContent struct {
	StartTime time.Time `json:"start_time"` // 开始时间
	EndTime   time.Time `json:"end_time"`   // 结束时间
	EndReason int8      `json:"end_reason"` // 结束原因 0-发起方挂断，1-接收方挂断 2-网络原因挂断 3-未打通
}

// VideoCallMsgContent
// @Description: 视频通话消息内容
// @Author Oberl-Fitzgerald 2024-06-14 18:56:54
type VideoCallMsgContent struct {
	StartTime time.Time `json:"start_time"` // 开始时间
	EndTime   time.Time `json:"end_time"`   // 结束时间
	EndReason int8      `json:"end_reason"` // 结束原因 0-发起方挂断，1-接收方挂断 2-网络原因挂断 3-未打通
}

// WithdrawMsgContent
// @Description: 撤回消息内容
// @Author Oberl-Fitzgerald 2024-06-14 18:57:38
type WithdrawMsgContent struct {
	OriginMsg *MsgContent `json:"_"`          // 原消息
	PromptMsg string      `json:"prompt_msg"` // 撤回的提示消息
}

// ReplyMsgContent
// @Description: 回复消息内容
// @Author Oberl-Fitzgerald 2024-06-14 18:59:29
type ReplyMsgContent struct {
	MsgID      uint        `json:"msg_id"`  // 消息ID
	Content    string      `json:"content"` // 回复内容
	MsgContent *MsgContent `json:"msg_content"`
}

// QuoteMsgContent
// @Description: 引用消息内容
// @Author Oberl-Fitzgerald 2024-06-14 19:03:51
type QuoteMsgContent struct {
	MsgID      uint        `json:"msg_id"`  // 消息ID
	Content    string      `json:"content"` // 回复内容
	MsgContent *MsgContent `json:"msg_content"`
}

// ForwardMsgContent
// @Description: 转发消息内容
// @Author Oberl-Fitzgerald 2024-06-14 19:04:04
type ForwardMsgContent struct {
	MsgID      uint        `json:"msg_id"` // 消息ID
	MsgContent *MsgContent `json:"msg_content"`
}
