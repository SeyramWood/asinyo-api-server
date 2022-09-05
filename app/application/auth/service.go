package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/domain/models"
)

type service struct {
	repo gateways.AuthRepo
}

// NewService is used to create a single instance of the service
func NewAuthService(repo gateways.AuthRepo) gateways.AuthService {
	return &service{
		repo: repo,
	}
}

func (s *service) Login(c *fiber.Ctx) error {
	authType := c.Get("Asinyo-Authorization-Type")
	var request models.User
	var merchantRequest models.UserMerchant

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
		return s.signinSupplierMerchant(c, merchantRequest)
	case "retailer":
		return s.signinRetailMerchant(c, merchantRequest)
	case "asinyo":
		return s.signinAdmin(c, request)
	default:
		return c.Status(fiber.StatusBadRequest).JSON(presenters.AuthErrorResponse("Bad request"))
	}
}

func (s *service) Logout(c *fiber.Ctx) error {
	c.Cookie(
		&fiber.Cookie{
			Name:     "remember",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HTTPOnly: true,
			SameSite: "none",
		},
	)
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"status": true,
		},
	)
}

func (s *service) FetchAuthUser(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)

	claims := user.Claims.(jwt.MapClaims)

	userType := claims["UserType"].(string)

	id := claims["Issuer"].(string)

	switch userType {
	case "customer":
		if user, err := s.repo.ReadCustomer(id, "id"); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"status": false,
				},
			)
		} else {
			return c.Status(fiber.StatusOK).JSON(presenters.AuthCustomerResponse(user))
		}
	case "agent":
		if user, err := s.repo.ReadAgent(id, "id"); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"status": false,
				},
			)
		} else {
			return c.Status(fiber.StatusOK).JSON(presenters.AuthAgentResponse(user))
		}
	case "supplier":
		if merchant, err := s.repo.ReadMerchant(id, "id"); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"status": false,
				},
			)
		} else {
			user, err := merchant.QuerySupplier().WithMerchant().Only(context.Background())
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(
					fiber.Map{
						"status": false,
					},
				)
			}
			return c.Status(fiber.StatusOK).JSON(
				presenters.AuthSupplierMerchantResponse(
					&presenters.AuthMerchant{
						ID:        user.Edges.Merchant.ID,
						Username:  user.Edges.Merchant.Username,
						LastName:  user.LastName,
						OtherName: user.OtherName,
					},
				),
			)

		}
	case "retailer":
		if merchant, err := s.repo.ReadMerchant(id, "id"); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"status": false,
				},
			)
		} else {
			user, err := merchant.QueryRetailer().WithMerchant().Only(context.Background())
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(
					fiber.Map{
						"status": false,
					},
				)
			}
			return c.Status(fiber.StatusOK).JSON(
				presenters.AuthRetailMerchantResponse(
					&presenters.AuthMerchant{
						ID:        user.Edges.Merchant.ID,
						Username:  user.Edges.Merchant.Username,
						LastName:  user.LastName,
						OtherName: user.OtherName,
					},
				),
			)
		}
	case "asinyo":
		if user, err := s.repo.ReadAdmin(id, "id"); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"status": false,
				},
			)
		} else {
			return c.Status(fiber.StatusOK).JSON(presenters.AuthAdminResponse(user))
		}
	default:
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"status": false,
			},
		)

	}

}

func (s *service) SendUserVerificationCode(username string) (string, error) {
	return "123456", nil
}

func (s *service) hashCheck(hash []byte, plain string) bool {
	if err := bcrypt.CompareHashAndPassword(hash, []byte(plain)); err != nil {
		return false
	}
	return true
}

// func (s *service) generateToken(id int, userType string) (interface{}, error) {

// 	claims := jwt.MapClaims{
// 		"Issuer":    strconv.Itoa(id),
// 		"IssuedAt":  jwt.NewNumericDate(time.Now()),
// 		"ExpiresAt": jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
// 		"UserType":  userType,
// 	}

// 	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	token, err := claim.SignedString([]byte(config.App().Key))

// 	if err != nil {
// 		return nil, err
// 	}

// 	return token, nil
// }

func (s *service) signinCustomer(c *fiber.Ctx, request models.User) error {

	if user, err := s.repo.ReadCustomer(request.Username, "username"); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(presenters.AuthErrorResponse("Bad credentials!"))
	} else {
		if s.hashCheck(user.Password, request.Password) {
			return c.Status(fiber.StatusOK).JSON(presenters.AuthCustomerResponse(user))
		}
	}
	return c.Status(fiber.StatusUnauthorized).JSON(presenters.AuthErrorResponse("Bad credentials"))
}

func (s *service) signinAgent(c *fiber.Ctx, request models.User) error {

	if user, err := s.repo.ReadAgent(request.Username, "username"); err != nil {
		fmt.Println(err)

		return c.Status(fiber.StatusNotFound).JSON(presenters.AuthErrorResponse("Bad credentials!"))
	} else {
		if s.hashCheck(user.Password, request.Password) {
			return c.Status(fiber.StatusOK).JSON(presenters.AuthAgentResponse(user))
		}
	}
	return c.Status(fiber.StatusUnauthorized).JSON(presenters.AuthErrorResponse("Bad credentials"))
}
func (s *service) signinSupplierMerchant(c *fiber.Ctx, request models.UserMerchant) error {

	if user, err := s.repo.ReadMerchant(request.Username, "username"); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(presenters.AuthErrorResponse("Bad credentials!"))
	} else {
		if s.hashCheck(user.Password, request.Password) {
			merchant, err := user.QuerySupplier().WithMerchant().Only(context.Background())
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(
					fiber.Map{
						"status": false,
					},
				)
			}
			return c.Status(fiber.StatusOK).JSON(
				presenters.AuthSupplierMerchantResponse(
					&presenters.AuthMerchant{
						ID:         merchant.Edges.Merchant.ID,
						Username:   merchant.Edges.Merchant.Username,
						LastName:   merchant.LastName,
						OtherName:  merchant.OtherName,
						Phone:      merchant.Phone,
						OtherPhone: *merchant.OtherPhone,
					},
				),
			)
		}
	}
	return c.Status(fiber.StatusUnauthorized).JSON(presenters.AuthErrorResponse("Bad credentials"))
}
func (s *service) signinRetailMerchant(c *fiber.Ctx, request models.UserMerchant) error {

	if user, err := s.repo.ReadMerchant(request.Username, "username"); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(presenters.AuthErrorResponse("Bad credentials!"))
	} else {
		if s.hashCheck(user.Password, request.Password) {
			merchant, err := user.QueryRetailer().WithMerchant().Only(context.Background())
			if err != nil {

				return c.Status(fiber.StatusInternalServerError).JSON(
					fiber.Map{
						"status": false,
					},
				)
			}
			return c.Status(fiber.StatusOK).JSON(
				presenters.AuthRetailMerchantResponse(
					&presenters.AuthMerchant{
						ID:         merchant.Edges.Merchant.ID,
						Username:   merchant.Edges.Merchant.Username,
						LastName:   merchant.LastName,
						OtherName:  merchant.OtherName,
						Phone:      merchant.Phone,
						OtherPhone: *merchant.OtherPhone,
					},
				),
			)
		}
	}
	return c.Status(fiber.StatusUnauthorized).JSON(presenters.AuthErrorResponse("Bad credentials"))
}
func (s *service) signinAdmin(c *fiber.Ctx, request models.User) error {

	if user, err := s.repo.ReadAdmin(request.Username, "username"); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(presenters.AuthErrorResponse("Bad credentials!"))
	} else {
		if s.hashCheck(user.Password, request.Password) {
			return c.Status(fiber.StatusOK).JSON(presenters.AuthAdminResponse(user))
		}
	}
	return c.Status(fiber.StatusUnauthorized).JSON(presenters.AuthErrorResponse("Bad credentials"))
}
