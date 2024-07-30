package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	FirstName string
	Email     string
	Password  string
	Plans     any
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
