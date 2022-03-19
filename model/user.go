package model

import (
	"time"
)

type Users struct {
	UUID      string    `gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password" gorm:"type:varchar(100)"`
	UpdatedAt time.Time `json:"updated_at"` // 신규 계정인 경우에는 계정 생성 당시를 기록함.
	CreatedAt time.Time `json:"created_at"`
}
