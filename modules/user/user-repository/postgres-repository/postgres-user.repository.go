package repositories

import (
	"trippy/db"
	"trippy/db/models"
	"trippy/modules/user/domain"
	"trippy/modules/user/mappers"

	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	database *gorm.DB
}

func NewPostgresUserRepository() *PostgresUserRepository {
	return &PostgresUserRepository{
		database: db.GetDB(),
	}
}

func (r *PostgresUserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	persistanceUser := mappers.ToPersistnaceUser(user)

	if err := r.database.Create(persistanceUser).Error; err != nil {
		return nil, err
	}

	return mappers.ToDomainUser(persistanceUser), nil
}
func (r *PostgresUserRepository) GetUserByEmail(email string) (*domain.User, error) {
	user := &models.User{}
	if err := r.database.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	return mappers.ToDomainUser(user), nil
}
func (r *PostgresUserRepository) GetUsers() ([]domain.User, error) {
	var users []models.User
	if err := r.database.Find(&users).Error; err != nil {
		return nil, err
	}
	var domainUsers []domain.User
	for _, user := range users {
		domainUser := mappers.ToDomainUser(&user)
		domainUsers = append(domainUsers, *domainUser)
	}
	return domainUsers, nil
}
