package token

import (
	"fmt"
	"prism-auth-service/pkg/webToken"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
)

func ValidateToken(c *fiber.Ctx) error {
	tokenData := c.FormValue("t")

	jwtToken, err := jwt.ParseWithClaims(tokenData, &webToken.Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return nil, fmt.Errorf("invalid token data")
	})

	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if claims, ok := jwtToken.Claims.(*webToken.Claims); ok && jwtToken.Valid {
		return c.JSON(fiber.Map{
			"Status": "success",
			"msg":    "Token valid",
			"data":   claims.Username,
		})
	}

	return c.SendStatus(fiber.StatusUnauthorized)
}

func RefreshToken(c *fiber.Ctx) error {

	return c.SendString("test2")
}
