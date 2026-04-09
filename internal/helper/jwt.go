package helper

import (
	"time"
	"user-api/internal/config"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret []byte

func InitJWT(c config.Config) {
	jwtSecret = []byte(c.Jwt.Secret)
}

func GetSecret() []byte {
	return jwtSecret
}

func GenerateToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
