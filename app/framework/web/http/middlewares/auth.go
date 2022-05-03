package middleware

import (
	"fmt"

	"github.com/SeyramWood/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("token")

		token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.App().Key), nil
		})
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthenticated",
			})
		}
		claims := token.Claims.(*jwt.RegisteredClaims)
		fmt.Println(claims.Issuer)
		return c.Next()
	}
}
