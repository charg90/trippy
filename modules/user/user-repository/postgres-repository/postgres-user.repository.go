package repositories

import (
	"trippy/db"
	"trippy/db/models"

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

func (r *PostgresUserRepository) CreateUser(user *models.User) (*models.User, error) {
	if err := r.database.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func (r *PostgresUserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	if err := r.database.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}