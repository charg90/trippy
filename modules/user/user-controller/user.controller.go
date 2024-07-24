package controller

import (
	"trippy/modules/user/dto"
	services "trippy/modules/user/user.service"
	"trippy/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	UserService services.UserService
	validator   *validator.Validate
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		UserService: userService,
		validator:  validator.New(),
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user dto.CreateUserDto
	
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	if err := c.validator.Struct(user); err != nil {
		utils.HandleValidationError(err, ctx)
		return
	}	


	createdUser, err := c.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	createUserResponse := dto.UserResponseDto{
		ID:    createdUser.ID,
		Name: createdUser.FirstName,
		Email: createdUser.Email,
	}

	ctx.JSON(201, createUserResponse)
}

func (c *UserController) Login(ctx *gin.Context) {
	var user dto.LoginDto

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.validator.Struct(user); err != nil {
		utils.HandleValidationError(err, ctx)
		return
	}


	token, err := c.UserService.Login(&user)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"User": token})
}

func (c *UserController) GetUsers(ctx *gin.Context) {
	users, err := c.UserService.GetUsers()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"users": users})
}