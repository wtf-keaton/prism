package webToken

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type Claims struct {
	jwt.StandardClaims

	Username string `json:"username"`
}

func Generate(username string) string {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(3600)),
			IssuedAt:  jwt.At(time.Now()),
		},
		Username: username,
	})

	generated, _ := jwtToken.SignedString([]byte("pipapupabubu"))

	return generated
}

func ValidateToken(jwtData string) (string, error) {
	jwtToken, err := jwt.ParseWithClaims(jwtData, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return nil, fmt.Errorf("invalid token data")
	})

	if err != nil {
		return "", errors.New(err.Error())
	}

	if claims, ok := jwtToken.Claims.(*Claims); ok && jwtToken.Valid {
		return claims.Username, nil
	}

	return "", errors.New("failed to verify token")

}
