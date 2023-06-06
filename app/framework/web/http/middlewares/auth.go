package middleware

import (
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/app_cache"
	"github.com/SeyramWood/app/application/auth"
	"github.com/SeyramWood/app/application/notification"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/config"
	jwtConf "github.com/SeyramWood/pkg/jwt"
)

const (
	REFRESH_TOKEN_EXPIRY = time.Minute * 120
	ACCESS_TOKEN_EXPIRY  = time.Minute * 60
)

type Middleware struct {
	authSrv  gateways.AuthService
	JWT      *jwtConf.JWT
	cache    *app_cache.AppCache
	userType map[string]string
}

func NewMiddleware(db *database.Adapter, noti notification.NotificationService, jwtConf *jwtConf.JWT, appCache *app_cache.AppCache) *Middleware {
	service := auth.NewAuthService(auth.NewAuthRepo(db), noti, jwtConf, appCache)
	return &Middleware{
		authSrv: service,
		JWT:     jwtConf,
		cache:   appCache,
		userType: map[string]string{
			"business":   "customer",
			"individual": "customer",
			"retailer":   "merchant",
			"supplier":   "merchant",
			"agent":      "agent",
			"asinyo":     "asinyo",
		},
	}
}

func (m *Middleware) IsAuthenticated() fiber.Handler {

	return func(c *fiber.Ctx) error {
		token := c.Cookies("__token")
		if ok := m.cache.Exist(token); !ok {
			__refresh := c.Cookies("__refresh")
			newTokens, err := m.authSrv.GenerateNewTokens(__refresh)
			if err != nil {
				log.Println(err)
				return c.Status(fiber.StatusForbidden).JSON(presenters.AuthErrorResponse("Unauthorized"))
			}
			userSession, ok := newTokens["session"].(*presenters.AuthSession)
			if !ok {
				log.Println(err)
				return c.Status(fiber.StatusForbidden).JSON(presenters.AuthErrorResponse("Unauthorized"))
			}
			if newTokens["token"].(string) != "" && newTokens["refresh"].(string) != "" {
				log.Println(newTokens["token"].(string), newTokens["refresh"].(string), "cookies")
			}

			c.Cookie(&fiber.Cookie{
				Name:     "userType",
				Value:    m.userType[userSession.UserType],
				Expires:  time.Now().Add(REFRESH_TOKEN_EXPIRY),
				Secure:   true,
				HTTPOnly: true,
			})
			c.Cookie(&fiber.Cookie{
				Name:     "__token",
				Value:    newTokens["token"].(string),
				Expires:  time.Now().Add(ACCESS_TOKEN_EXPIRY),
				Secure:   true,
				HTTPOnly: true,
			})
			c.Cookie(&fiber.Cookie{
				Name:     "__refresh",
				Value:    newTokens["refresh"].(string),
				Expires:  time.Now().Add(REFRESH_TOKEN_EXPIRY),
				Secure:   true,
				HTTPOnly: true,
			})
			return c.Next()

		}
		return c.Next()
	}
}

func (m *Middleware) IsAuthorized() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token, err := m.bearerToken(c, "Authorization")
		if err != nil {
			log.Println("request failed API Token authentication\nerror: ", err, " - Remote IP: ", c.IP())
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "403 - Forbidden"})
		}
		if _, ok := m.isValidToken(token); !ok {
			log.Println("no matching API token found - Remote IP: ", c.IP())
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "403 - Forbidden"})
		}
		return c.Next()
	}
}

// bearerToken extracts the content from the header, striping the Bearer prefix
func (m *Middleware) bearerToken(c *fiber.Ctx, header string) (string, error) {
	rawToken := c.Get(header)
	pieces := strings.SplitN(rawToken, " ", 2)
	if len(pieces) < 2 {
		return "", fmt.Errorf("token with incorrect bearer format")
	}
	token := strings.TrimSpace(pieces[1])
	return token, nil
}

// apiKeyIsValid checks if the given API key is valid and returns the principal if it is.
func (m *Middleware) isValidToken(token string) (string, bool) {
	secrete := fmt.Sprintf("%x", sha256.Sum256([]byte(config.App().Key)))
	contentEqual := subtle.ConstantTimeCompare([]byte(token), []byte(secrete)) == 1
	if contentEqual {
		return token, true
	}
	return "", false
}
