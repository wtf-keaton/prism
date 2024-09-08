package webToken

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type Claims struct {
	jwt.StandardClaims

	Username string `json:"username"`
}

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func Generate(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.At(time.Now()),
		},
		Username: username,
	})

	return token.SignedString(secretKey)
}

func ValidateToken(jwtData string) (string, error) {
	token, err := jwt.ParseWithClaims(jwtData, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return "", fmt.Errorf("failed to parse token: %v", err)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.Username, nil
	}

	return "", errors.New("invalid token")
}
