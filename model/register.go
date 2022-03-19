package model

import (
	"time"

	"github.com/google/uuid"
)

type Register struct {
	UUID      uuid.UUID
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
