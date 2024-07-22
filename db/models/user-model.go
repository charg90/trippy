package models

import "time"

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt *time.Time     `gorm:"index" json:"deletedAt"`
	FirstName string `json:"firstName"` 
    Password  string `json:"password"` 
	Email     string `json:"email"`
}