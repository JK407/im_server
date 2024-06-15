package ctype

import (
	"database/sql/driver"
	"encoding/json"
)

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

// Scan
// @Description Scan 从数据库中读取的数据
// @Author Oberl-Fitzgerald 2024-06-14 20:05:32
// @Param  value interface{}
// @Return error
func (v *VerifyQuestion) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), v)
}

// Value
// @Description Value 入库的数据
// @Author Oberl-Fitzgerald 2024-06-14 20:05:35
// @Return value
// @Return err
func (v VerifyQuestion) Value() (driver.Value, error) {
	b, err := json.Marshal(v)
	return string(b), err
}
