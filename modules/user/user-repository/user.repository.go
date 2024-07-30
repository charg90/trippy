package repositories

import "trippy/modules/user/domain"

type UserRepository interface {
	CreateUser(user *domain.User) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	GetUsers() ([]domain.User, error)
}
