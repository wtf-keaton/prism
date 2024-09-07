package webToken

import (
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
