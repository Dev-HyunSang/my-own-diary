package model

import (
	"time"

	"gorm.io/gorm"
)

type Diary struct {
	gorm.Model
	DiaryUUID string    `json:"diary_uuid" gorm:"primaryKey;"`
	UserUUID  string    `json:"user_uuid"`
	Group     string    `json:"diary_group"`
	Content   string    `json:"content"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	RevisedAt time.Time `json:"revised_at"`
}
