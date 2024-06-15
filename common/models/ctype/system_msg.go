package ctype

import (
	"database/sql/driver"
	"encoding/json"
)

// SystemMsg
// @Description: 系统提示
// @Author Oberl-Fitzgerald 2024-06-14 19:06:09
type SystemMsg struct {
	Type int8 `json:"type"` // 违规类型 1-色情，2-广告，3-政治，4-诈骗，5-其他
}

// Scan
// @Description Scan 从数据库中读取的数据
// @Author Oberl-Fitzgerald 2024-06-14 20:04:04
// @Param  value interface{}
// @Return error
func (s *SystemMsg) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), s)
}

// Value
// @Description Value 入库的数据
// @Author Oberl-Fitzgerald 2024-06-14 20:04:08
// @Return value
// @Return err
func (s SystemMsg) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return string(b), err
}
