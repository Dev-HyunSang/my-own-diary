package model

import (
	"time"

	"gorm.io/gorm"
)

type Diary struct {
	gorm.Model
	DiaryUUID string    `json:"diary_uuid" gorm:"primaryKey;"`
	UserUUID  string    `json:"user_uuid"` // 생성한 유저의 UUID를 기록합니다.
	Group     string    `json:"group"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	RevisedAt time.Time `json:"revised_at"`
}
