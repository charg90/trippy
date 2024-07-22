package services

import (
	"fmt"
	"trippy/db/models"
	"trippy/modules/user/dto"
	repositories "trippy/modules/user/user-repository"
	"trippy/utils"
)

type UserService interface {
	CreateUser(user *dto.CreateUserDto) (*models.User, error)
	Login(user *dto.LoginDto) (*dto.LoginResponseDto, error)
}

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (s *UserServiceImpl) CreateUser(userDto *dto.CreateUserDto) (*models.User, error) {
	
	hashedPassword, err := utils.HashPassword(userDto.Password)
	if err != nil {
		 fmt.Println("Error hashing password: ", err)
		return nil, err
	}

	userDto.Password = hashedPassword

	user:= &models.User{
		FirstName: userDto.FirstName,
		Password:  userDto.Password,
		Email:     userDto.Email,
	}
	return s.UserRepository.CreateUser(user)
}

func (s *UserServiceImpl) Login(userDto *dto.LoginDto) (*dto.LoginResponseDto, error) {
	user, err := s.UserRepository.GetUserByEmail(userDto.Email)
	if err != nil {
		return nil, err
	}
	
	err = utils.ComparePassword(userDto.Password, user.Password)
	if err != nil {
		return nil, err
	}
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return nil, err
	}
	return &dto.LoginResponseDto{
		Token:     token,
		Email:     user.Email,
		FirstName: user.FirstName,
	}, nil
}