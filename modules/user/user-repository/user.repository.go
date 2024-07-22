package repositories

import "trippy/db/models"


type UserRepository interface{
 CreateUser(user *models.User) (*models.User, error)
 GetUserByEmail(email string) (*models.User, error)
}