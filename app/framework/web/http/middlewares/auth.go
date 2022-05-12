package middleware

import (
	"fmt"

	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/config"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func Auth2() fiber.Handler {
	return func(c *fiber.Ctx) error {

		cookie := c.Cookies("asinyo_remember")

		token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.App().Key), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(presenters.AuthErrorResponse("Unauthenticated"))
		}

		claims := (*token.Claims.(*jwt.MapClaims))

		fmt.Println(claims["UserType"])

		return c.Next()
	}
}

func Auth() func(c *fiber.Ctx) error {

	return jwtware.New(jwtware.Config{

		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthenticated"})
		},
		SigningKey: []byte(config.App().Key),
	})

}
