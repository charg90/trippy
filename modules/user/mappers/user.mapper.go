package mappers

import (
	"trippy/db/models"
	"trippy/modules/user/domain"
)

func ToDomainUser(user *models.User) *domain.User {
	return &domain.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}

func ToPersistnaceUser(user *domain.User) *models.User {
	return &models.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		Email:     user.Email,
		Password:  user.Password,
	}
}
