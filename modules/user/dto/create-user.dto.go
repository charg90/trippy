package dto

import "github.com/google/uuid"

type CreateUserDto struct {
	FirstName string `json:"firstName" validate:"required,min=2,max=100"`
	Password  string `json:"password" validate:"required,min=8,max=100"`
	Email     string `json:"email" validate:"required,email"`
}

type UserResponseDto struct {
    ID    uuid.UUID    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}