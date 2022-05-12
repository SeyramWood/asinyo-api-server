package auth

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo gateways.AuthRepo
}

//NewService is used to create a single instance of the service
func NewAuthService(repo gateways.AuthRepo) gateways.AuthService {
	return &service{
		repo: repo,
	}
}

func (s *service) Login(c *fiber.Ctx) error {
	authType := c.Get("Asinyo-Authorization-Type")
	var request models.User
	var merchantRequest models.UserMerchant
	fmt.Println(authType)
	if authType == "supplier" || authType == "retailer" {
		err := c.BodyParser(&merchantRequest)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.AuthErrorResponse("Bad request"))
		}

	} else {
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.AuthErrorResponse("Bad request"))
		}
	}

	switch authType {
	case "customer":
		return s.signinCustomer(c, request)
	case "agent":
		return s.signinAgent(c, request)
	case "supplier":
		return s.signinSupplerMerchant(c, merchantRequest)
	case "retailer":
		return s.signinRetailMerchant(c, merchantRequest)
	case "asinyo":
		return s.signinAdmin(c, request)
	default:
		return c.Status(fiber.StatusBadRequest).JSON(presenters.AuthErrorResponse("Bad request"))
	}
}

func (s *service) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "remember",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		SameSite: "none",
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
	})
}

func (s *service) FetcAuthUser(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)

	claims := user.Claims.(jwt.MapClaims)

	userType := claims["UserType"].(string)

	id := claims["Issuer"].(string)

	switch userType {
	case "customer":
		if user, err := s.repo.ReadCustomer(id, "id"); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": false,
			})
		} else {
			return c.Status(fiber.StatusOK).JSON(presenters.AuthCustomerResponse(user))
		}
	case "agent":
		if user, err := s.repo.ReadAgent(id, "id"); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": false,
			})
		} else {
			return c.Status(fiber.StatusOK).JSON(presenters.AuthAgentResponse(user))
		}
	case "supplier":
		if merchant, err := s.repo.ReadMerchant(id, "id"); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": false,
			})
		} else {
			user, err := merchant.QuerySupplier().WithMerchant().Only(context.Background())
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status": false,
				})
			}
			return c.Status(fiber.StatusOK).JSON(presenters.AuthSupplierMerchantResponse(&presenters.AuthSupplierMerchant{
				ID:        user.Edges.Merchant.ID,
				Username:  user.Edges.Merchant.Username,
				LastName:  user.LastName,
				OtherName: user.OtherName,
			}))

		}
	case "retailer":
		if merchant, err := s.repo.ReadMerchant(id, "id"); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": false,
			})
		} else {
			user, err := merchant.QueryRetailer().WithMerchant().Only(context.Background())
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status": false,
				})
			}
			return c.Status(fiber.StatusOK).JSON(presenters.AuthRetailMerchantResponse(&presenters.AuthRetailMerchant{
				ID:        user.Edges.Merchant.ID,
				Username:  user.Edges.Merchant.Username,
				LastName:  user.LastName,
				OtherName: user.OtherName,
			}))
		}
	case "asinyo":
		if user, err := s.repo.ReadAdmin(id, "id"); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": false,
			})
		} else {
			return c.Status(fiber.StatusOK).JSON(presenters.AuthAdminResponse(user))
		}
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
		})

	}

}

func (s *service) hashCheck(hash []byte, plain string) bool {
	if err := bcrypt.CompareHashAndPassword(hash, []byte(plain)); err != nil {
		return false
	}
	return true
}

func (s *service) geterateToken(id int, userType string) (interface{}, error) {

	claims := jwt.MapClaims{
		"Issuer":    strconv.Itoa(id),
		"IssuedAt":  jwt.NewNumericDate(time.Now()),
		"ExpiresAt": jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		"UserType":  userType,
	}

	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := claim.SignedString([]byte(config.App().Key))

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *service) signinCustomer(c *fiber.Ctx, request models.User) error {

	if user, err := s.repo.ReadCustomer(request.Username, "username"); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(presenters.AuthErrorResponse("Bad credentials!"))
	} else {
		if s.hashCheck(user.Password, request.Password) {
			if token, err := s.geterateToken(user.ID, "customer"); err != nil {
				return c.SendStatus(fiber.StatusInternalServerError)
			} else {
				return c.Status(fiber.StatusOK).JSON(fiber.Map{config.App().TokenName: token.(string)})
			}
		}
	}
	return c.Status(fiber.StatusUnauthorized).JSON(presenters.AuthErrorResponse("Bad credentials"))
}

func (s *service) signinAgent(c *fiber.Ctx, request models.User) error {

	if user, err := s.repo.ReadAgent(request.Username, "username"); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(presenters.AuthErrorResponse("Bad credentials!"))
	} else {
		if s.hashCheck(user.Password, request.Password) {
			if token, err := s.geterateToken(user.ID, "agent"); err != nil {
				return c.SendStatus(fiber.StatusInternalServerError)
			} else {
				return c.Status(fiber.StatusOK).JSON(fiber.Map{config.App().TokenName: token.(string)})
			}
		}
	}
	return c.Status(fiber.StatusUnauthorized).JSON(presenters.AuthErrorResponse("Bad credentials"))
}
func (s *service) signinSupplerMerchant(c *fiber.Ctx, request models.UserMerchant) error {

	if user, err := s.repo.ReadMerchant(request.Username, "username"); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(presenters.AuthErrorResponse("Bad credentials!"))
	} else {
		fmt.Println(user)
		if s.hashCheck(user.Password, request.Password) {
			if token, err := s.geterateToken(user.ID, "supplier"); err != nil {
				return c.SendStatus(fiber.StatusInternalServerError)
			} else {
				return c.Status(fiber.StatusOK).JSON(fiber.Map{config.App().TokenName: token.(string)})
			}
		}
	}
	return c.Status(fiber.StatusUnauthorized).JSON(presenters.AuthErrorResponse("Bad credentials"))
}
func (s *service) signinRetailMerchant(c *fiber.Ctx, request models.UserMerchant) error {

	if user, err := s.repo.ReadMerchant(request.Username, "username"); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(presenters.AuthErrorResponse("Bad credentials!"))
	} else {
		if s.hashCheck(user.Password, request.Password) {
			if token, err := s.geterateToken(user.ID, "retailer"); err != nil {
				return c.SendStatus(fiber.StatusInternalServerError)
			} else {
				return c.Status(fiber.StatusOK).JSON(fiber.Map{config.App().TokenName: token.(string)})
			}
		}
	}
	return c.Status(fiber.StatusUnauthorized).JSON(presenters.AuthErrorResponse("Bad credentials"))
}
func (s *service) signinAdmin(c *fiber.Ctx, request models.User) error {

	if user, err := s.repo.ReadAdmin(request.Username, "username"); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(presenters.AuthErrorResponse("Bad credentials!"))
	} else {
		if s.hashCheck(user.Password, request.Password) {
			if token, err := s.geterateToken(user.ID, "asinyo"); err != nil {
				return c.SendStatus(fiber.StatusInternalServerError)
			} else {
				return c.Status(fiber.StatusOK).JSON(fiber.Map{config.App().TokenName: token.(string)})
			}
		}
	}
	return c.Status(fiber.StatusUnauthorized).JSON(presenters.AuthErrorResponse("Bad credentials"))
}
