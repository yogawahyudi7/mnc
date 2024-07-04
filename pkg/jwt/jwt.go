package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/yogawahyudi7/mnc/config"
)

func GenerateToken(config *config.Server, id string, tokenType string, tokenDuration string) (string, error) {

	timeDuration, _ := time.ParseDuration(tokenDuration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        id,
		"tokenType": tokenType,
		"exp":       time.Now().Add(timeDuration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(config *config.Server, tokenString string) (jwt.MapClaims, error) {

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}
