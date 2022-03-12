package model

import "github.com/google/uuid"

type Register struct {
	UUID     uuid.UUID
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
