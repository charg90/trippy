package dto

import (
	"time"

	"github.com/google/uuid"
)

type ResponseGetAllUsers struct {
	ID        uuid.UUID  `json:"id"`
	FirstName string     `json:"firstName"`
	Email     string     `json:"email"`
	Plans     any        `json:"plans"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
