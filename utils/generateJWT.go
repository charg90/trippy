package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)
type Claims struct {
    UserID uuid.UUID `json:"user_id"`
    jwt.StandardClaims
}

func GenerateJWT(id uuid.UUID) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func GenerateRefreshToken(id uuid.UUID) (string, error) {
    claims := jwt.MapClaims{}
    claims["user_id"] = id
    claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix() // 1 week expiry
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(os.Getenv("JWT_REFRESH_SECRET")))
}

func ValidateToken(tokenString string, secret []byte) (*Claims, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return secret, nil
    })

    if err != nil || !token.Valid {
        return nil, err
    }

    return claims, nil
}