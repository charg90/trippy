package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt *time.Time     `gorm:"index" json:"deletedAt"`
	FirstName string `json:"firstName"` 
    Password  string `json:"password"` 
	Email     string `json:"email"`
}
