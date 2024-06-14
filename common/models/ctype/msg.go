package ctype

import (
	"encoding/json"
	"time"
)

// Msg
// @Description: 消息内容
// @Author Oberl-Fitzgerald 2024-06-14 18:53:41
type Msg struct {
	Type         int8          `json:"type"`           // 消息类型：1-文本，2-图片，3-视频，4-文件，5-语音， 6-语音通话， 7-视频通话， 8-撤回消息， 9-回复消息， 10-引用消息， 11-转发消息 12-艾特消息
	Content      *string       `json:"content"`        // 文本内容 为1时使用
	ImageMsg     *ImageMsg     `json:"image_msg"`      // 图片内容 为2时使用
	VideoMsg     *VideoMsg     `json:"video_msg"`      // 视频内容 为3时使用
	FileMsg      *FileMsg      `json:"file_msg"`       // 文件内容 为4时使用
	VoiceMsg     *VoiceMsg     `json:"voice_msg"`      // 语音内容 为5时使用
	VoiceCallMsg *VoiceCallMsg `json:"voice_call_msg"` // 语音通话内容 为6时使用
	VideoCallMsg *VideoCallMsg `json:"video_call_msg"` // 视频通话内容 为7时使用
	WithdrawMsg  *WithdrawMsg  `json:"withdraw_msg"`   // 撤回消息内容 为8时使用
	ReplyMsg     *ReplyMsg     `json:"reply_msg"`      // 回复消息内容 为9时使用
	QuoteMsg     *QuoteMsg     `json:"quote_msg"`      // 引用消息内容 为10时使用
	ForwardMsg   *ForwardMsg   `json:"forward_msg"`    // 转发消息内容 为11时使用
	AtMsg        *AtMsg        `json:"at_msg"`         // 艾特消息内容 为12时使用 群聊才有
}

// Scan
// @Description Scan 从数据库中读取的数据
// @Author Oberl-Fitzgerald 2024-06-14 20:03:24
// @Param  value interface{}
// @Return error
func (m *Msg) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), m)
}

// Value
// @Description Value 入库的数据
// @Author Oberl-Fitzgerald 2024-06-14 20:02:58
// @Return value
// @Return err
func (m Msg) Value() (value interface{}, err error) {
	b, err := json.Marshal(m)
	return string(b), err
}

// ImageMsg
// @Description: 图片消息内容
// @Author Oberl-Fitzgerald 2024-06-14 18:53:45
type ImageMsg struct {
	Title string `json:"title"` // 标题
	URL   string `json:"url"`   // 地址
}

// VideoMsg
// @Description: 视频消息内容
// @Author Oberl-Fitzgerald 2024-06-14 18:53:47
type VideoMsg struct {
	Title string `json:"title"` // 标题
	URL   string `json:"url"`   // 地址
	Time  int    `json:"time"`  // 时长 单位秒
}

// FileMsg
// @Description: 文件消息内容
// @Author Oberl-Fitzgerald 2024-06-14 18:53:51
type FileMsg struct {
	Title string `json:"title"` // 标题
	URL   string `json:"url"`   // 地址
	Size  int64  `json:"size"`  // 大小 单位字节
	Type  string `json:"type"`  // 文件类型
}

// VoiceMsg
// @Description: 语音消息内容
// @Author Oberl-Fitzgerald 2024-06-14 18:53:53
type VoiceMsg struct {
	URL  string `json:"url"`  // 地址
	Time int    `json:"time"` // 时长 单位秒
}

// VoiceCallMsg
// @Description: 语音通话消息内容
// @Author Oberl-Fitzgerald 2024-06-14 18:54:12
type VoiceCallMsg struct {
	StartTime time.Time `json:"start_time"` // 开始时间
	EndTime   time.Time `json:"end_time"`   // 结束时间
	EndReason int8      `json:"end_reason"` // 结束原因 0-发起方挂断，1-接收方挂断 2-网络原因挂断 3-未打通
}

// VideoCallMsg
// @Description: 视频通话消息内容
// @Author Oberl-Fitzgerald 2024-06-14 18:56:54
type VideoCallMsg struct {
	StartTime time.Time `json:"start_time"` // 开始时间
	EndTime   time.Time `json:"end_time"`   // 结束时间
	EndReason int8      `json:"end_reason"` // 结束原因 0-发起方挂断，1-接收方挂断 2-网络原因挂断 3-未打通
}

// WithdrawMsg
// @Description: 撤回消息内容
// @Author Oberl-Fitzgerald 2024-06-14 18:57:38
type WithdrawMsg struct {
	OriginMsg *Msg   `json:"_"`          // 原消息
	PromptMsg string `json:"prompt_msg"` // 撤回的提示消息
}

// ReplyMsg
// @Description: 回复消息内容
// @Author Oberl-Fitzgerald 2024-06-14 18:59:29
type ReplyMsg struct {
	MsgID   uint   `json:"msg_id"`  // 消息ID
	Content string `json:"content"` // 回复内容
	Msg     *Msg   `json:"msg"`
}

// QuoteMsg
// @Description: 引用消息内容
// @Author Oberl-Fitzgerald 2024-06-14 19:03:51
type QuoteMsg struct {
	MsgID   uint   `json:"msg_id"`  // 消息ID
	Content string `json:"content"` // 回复内容
	Msg     *Msg   `json:"msg"`
}

// ForwardMsg
// @Description: 转发消息内容
// @Author Oberl-Fitzgerald 2024-06-14 19:04:04
type ForwardMsg struct {
	MsgID uint `json:"msg_id"` // 消息ID
	Msg   *Msg `json:"msg"`
}

// AtMsg
// @Description: 艾特消息内容
// @Author Oberl-Fitzgerald 2024-06-14 19:53:54
type AtMsg struct {
	UserID  uint   `json:"user_id"` // 艾特的用户ID
	Content string `json:"content"` // 回复内容
	Msg     *Msg   `json:"msg"`
}
