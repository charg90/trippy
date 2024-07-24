package dto

type LoginDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponseDto struct {
	Token string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	Email string `json:"email"`
	FirstName string `json:"firstName"`

}