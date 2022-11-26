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
	"github.com/SeyramWood/app/application"
	"github.com/SeyramWood/app/application/sms"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/domain/services"
)

type service struct {
	repo gateways.AuthRepo
	sms  gateways.SMSService
	mail gateways.EmailService
}

func NewAuthService(repo gateways.AuthRepo, mail gateways.EmailService) gateways.AuthService {
	smsService := sms.NewSMSService()
	return &service{
		repo: repo,
		sms:  smsService,
		mail: mail,
	}
}

func (s *service) Login(c *fiber.Ctx) error {
	authType := c.Get("Asinyo-Authorization-Type")
	var request models.User
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presenters.AuthErrorResponse("Bad request"))
	}

	switch authType {
	case "customer":
		return s.signinCustomer(c, request)
	case "agent":
		return s.signinAgent(c, request)
	case "merchant":
		return s.signinMerchant(c, request)
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
			_, err := merchant.QuerySupplier().WithMerchant().Only(context.Background())
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(
					fiber.Map{
						"status": false,
					},
				)
			}
			// return c.Status(fiber.StatusOK).JSON(
			// 	presenters.AuthSupplierMerchantResponse(
			// 		&presenters.AuthMerchant{
			// 			ID:        user.Edges.Merchant.ID,
			// 			Username:  user.Edges.Merchant.Username,
			// 			LastName:  user.LastName,
			// 			OtherName: user.OtherName,
			// 		},
			// 	),
			// )

		}
	case "retailer":
		if merchant, err := s.repo.ReadMerchant(id, "id"); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"status": false,
				},
			)
		} else {
			_, err := merchant.QueryRetailer().WithMerchant().Only(context.Background())
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(
					fiber.Map{
						"status": false,
					},
				)
			}
			// return c.Status(fiber.StatusOK).JSON(
			// 	presenters.AuthRetailMerchantResponse(
			// 		&presenters.AuthMerchant{
			// 			ID:        user.Edges.Merchant.ID,
			// 			Username:  user.Edges.Merchant.Username,
			// 			LastName:  user.LastName,
			// 			OtherName: user.OtherName,
			// 		},
			// 	),
			// )
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
	return nil
}

func (s *service) UpdatePassword(id string, request any, userType string, isOTP bool) (bool, error) {
	data := request.(*models.ChangePassword)
	switch userType {
	case "customer":
		if user, err := s.repo.ReadCustomer(id, "id"); err != nil {
			return false, fmt.Errorf("no record found for %d", user.ID)
		} else {
			if s.hashCheck(user.Password, data.CurrentPassword) {
				return s.repo.UpdatePassword(user.ID, data.Password, userType, isOTP)
			}
		}
		return false, fmt.Errorf("current password do not match our records")
	case "agent":
		if user, err := s.repo.ReadAgent(id, "id"); err != nil {
			return false, fmt.Errorf("no record found for %d", user.ID)
		} else {
			if s.hashCheck(user.Password, data.CurrentPassword) {
				return s.repo.UpdatePassword(user.ID, data.Password, userType, isOTP)
			}
		}
		return false, fmt.Errorf("current password do not match our records")
	case "supplier", "retailer":
		if user, err := s.repo.ReadMerchant(id, "id"); err != nil {
			return false, fmt.Errorf("no record found for %d", user.ID)
		} else {
			if s.hashCheck(user.Password, data.CurrentPassword) {
				return s.repo.UpdatePassword(user.ID, data.Password, userType, isOTP)
			}
		}
		return false, fmt.Errorf("current password do not match our records")
	case "asinyo":
		if user, err := s.repo.ReadAdmin(id, "id"); err != nil {
			return false, fmt.Errorf("no record found for %d", user.ID)
		} else {
			if s.hashCheck(user.Password, data.CurrentPassword) {
				return s.repo.UpdatePassword(user.ID, data.Password, userType, isOTP)
			}
		}
		return false, fmt.Errorf("current password do not match our records")
	default:
		return false, fmt.Errorf("current password do not match our records")
	}
}

func (s *service) ResetPassword(request *models.ResetPassword, username, userType string) (bool, error) {
	switch userType {
	case "customer":
		if user, err := s.repo.ReadCustomer(username, "username"); err != nil {
			return false, fmt.Errorf("no record found for %s", user.Username)
		} else {
			return s.repo.ResetPassword(user.ID, request.NewPassword, userType)
		}

	case "agent":
		if user, err := s.repo.ReadAgent(username, "username"); err != nil {
			return false, fmt.Errorf("no record found for %s", user.Username)
		} else {
			return s.repo.ResetPassword(user.ID, request.NewPassword, userType)
		}
	case "merchant":
		if user, err := s.repo.ReadMerchant(username, "username"); err != nil {
			return false, fmt.Errorf("no record found for %s", user.Username)
		} else {
			return s.repo.ResetPassword(user.ID, request.NewPassword, userType)
		}
	case "asinyo":
		if user, err := s.repo.ReadAdmin(username, "username"); err != nil {
			return false, fmt.Errorf("no record found for %s", user.Username)
		} else {
			return s.repo.ResetPassword(user.ID, request.NewPassword, userType)
		}
	default:
		return false, fmt.Errorf("no record found")
	}
}

func (s *service) SendUserVerificationCode(username string) (string, error) {
	code, _ := application.GenerateOTP(6)
	msg := fmt.Sprintf(
		"Congratulations for your attempt to join Asinyo! Please enter the OTP Code to proceed with your sign up. %s",
		code,
	)
	if application.UsernameType(username, "phone") {
		_, err := s.sms.Send(
			&services.SMSPayload{
				Recipients: []string{username},
				Message:    msg,
			},
		)
		if err != nil {
			return "", err
		}
	}
	if application.UsernameType(username, "email") {

		s.mail.Send(
			&services.Message{
				To:      username,
				Subject: "ASINYO SIGN UP VERIFICATION",
				Data:    msg,
			},
		)
	}
	return code, nil
}
func (s *service) SendPasswordResetCode(username, userType string) (string, error) {
	if userType == "customer" {
		_, err := s.repo.ReadCustomer(username, "username")
		if err != nil {
			return "", nil
		}
	}
	if userType == "agent" {
		_, err := s.repo.ReadAgent(username, "username")
		if err != nil {
			return "", nil
		}
	}
	if userType == "merchant" {
		_, err := s.repo.ReadMerchant(username, "username")
		if err != nil {
			return "", nil
		}
	}
	if userType == "asinyo" {
		_, err := s.repo.ReadAdmin(username, "username")
		if err != nil {
			return "", nil
		}
	}

	code, _ := application.GenerateOTP(6)
	msg := fmt.Sprintf(
		"You are one step away to complete your password reset! Please enter the RESET Code to proceed. %s",
		code,
	)
	if application.UsernameType(username, "phone") {
		_, err := s.sms.Send(
			&services.SMSPayload{
				Recipients: []string{username},
				Message:    msg,
			},
		)
		if err != nil {
			return "", err
		}
	}
	if application.UsernameType(username, "email") {

		s.mail.Send(
			&services.Message{
				To:      username,
				Subject: "ASINYO PASSWORD RESET",
				Data:    msg,
			},
		)
	}

	return code, nil

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
		return c.Status(fiber.StatusNotFound).JSON(presenters.AuthErrorResponse("Bad credentials!"))
	} else {
		if s.hashCheck(user.Password, request.Password) {
			return c.Status(fiber.StatusOK).JSON(presenters.AuthAgentResponse(user))
		}
	}
	return c.Status(fiber.StatusUnauthorized).JSON(presenters.AuthErrorResponse("Bad credentials"))
}
func (s *service) signinMerchant(c *fiber.Ctx, request models.User) error {

	if user, err := s.repo.ReadMerchant(request.Username, "username"); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(presenters.AuthErrorResponse("Bad credentials!"))
	} else {
		if s.hashCheck(user.Password, request.Password) {
			return c.Status(fiber.StatusOK).JSON(presenters.AuthMerchantResponse(user))
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
