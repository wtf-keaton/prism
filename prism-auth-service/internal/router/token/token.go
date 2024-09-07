package token

import (
	"fmt"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
)

type Claims struct {
	jwt.StandardClaims

	Username string `json:"username"`
}

func ValidateToken(c *fiber.Ctx) error {
	tokenData := c.FormValue("t")

	jwtToken, err := jwt.ParseWithClaims(tokenData, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return nil, nil
	})

	if err != nil {
		return c.JSON(fiber.Map{
			"Status": "failed",
			"msg":    "Unexpected signing method",
		})
	}

	if claims, ok := jwtToken.Claims.(*Claims); ok && jwtToken.Valid {
		return c.JSON(fiber.Map{
			"Status": "success",
			"msg":    "Token valid",
			"data":   claims.Username,
		})
	}

	return c.JSON(fiber.Map{
		"Status": "failed",
		"msg":    "Failed to validate token",
	})
}

func RefreshToken(c *fiber.Ctx) error {

	return c.SendString("test2")
}
