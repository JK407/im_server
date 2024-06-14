package models

import "time"

type Models struct {
	ID        uint      `json:"id"`         // ID
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
}
