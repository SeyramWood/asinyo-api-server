package auth

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"golang.org/x/crypto/bcrypt"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application"
	"github.com/SeyramWood/app/application/notification"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/config"
	"github.com/SeyramWood/ent"
	jwtUtil "github.com/SeyramWood/pkg/jwt"
)

const (
	RefreshTokenExpiry = time.Minute * 120
	AccessTokenExpiry  = time.Minute * 60
)

type service struct {
	repo     gateways.AuthRepo
	noti     notification.NotificationService
	JWT      *jwtUtil.JWT
	cache    gateways.CacheService
	userType map[string]string
}

func NewAuthService(
	repo gateways.AuthRepo, noti notification.NotificationService, JWT *jwtUtil.JWT, cache gateways.CacheService,
) gateways.AuthService {

	return &service{
		repo:  repo,
		noti:  noti,
		JWT:   JWT,
		cache: cache,
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

func (s *service) Login(c *fiber.Ctx) error {
	authType := c.Get("Asinyo-Authorization-Type")
	var request models.User
	if err := c.BodyParser(&request); err != nil {
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
	token := c.Cookies("__token")
	if ok := s.cache.Exist(token); !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"status": false,
				"msg":    "Unauthorized",
			},
		)
	}
	_ = s.cache.Delete(token)
	c.Cookie(
		&fiber.Cookie{
			Name:    "__token",
			Value:   "",
			Expires: time.Now().Add(-AccessTokenExpiry),
		},
	)
	c.Cookie(
		&fiber.Cookie{
			Name:    "__refresh",
			Value:   "",
			Expires: time.Now().Add(-AccessTokenExpiry),
		},
	)
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"status": true,
		},
	)
}

func (s *service) RefreshToken(c *fiber.Ctx) error {
	__refresh := c.Cookies("__refresh")
	if ok := s.cache.Exist(__refresh); !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"status": false,
				"msg":    "Unauthorized",
			},
		)
	}
	newTokens, err := s.GenerateNewTokens(__refresh)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(presenters.AuthErrorResponse("Unauthorized"))
	}
	userSession, ok := newTokens["session"].(*presenters.AuthSession)
	if !ok {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(presenters.AuthErrorResponse("Unauthorized"))
	}
	err = s.cache.Set(newTokens["token"].(string), &userSession, AccessTokenExpiry)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(presenters.AuthErrorResponse(err))
	}
	err = s.cache.Set(newTokens["refresh"].(string), &userSession, RefreshTokenExpiry)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(presenters.AuthErrorResponse(err))
	}
	c.Cookie(
		&fiber.Cookie{
			Name:     "userType",
			Value:    s.userType[userSession.UserType],
			Expires:  time.Now().Add(RefreshTokenExpiry),
			Secure:   true,
			HTTPOnly: false,
		},
	)
	c.Cookie(
		&fiber.Cookie{
			Name:     "__token",
			Value:    newTokens["token"].(string),
			Expires:  time.Now().Add(AccessTokenExpiry),
			Secure:   true,
			HTTPOnly: false,
		},
	)
	c.Cookie(
		&fiber.Cookie{
			Name:     "__refresh",
			Value:    newTokens["refresh"].(string),
			Expires:  time.Now().Add(RefreshTokenExpiry),
			Secure:   true,
			HTTPOnly: false,
		},
	)
	// _ = s.cache.Delete(__refresh)
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"status":   true,
			"userType": s.userType[userSession.UserType],
			"token":    newTokens["token"].(string),
			"refresh":  newTokens["refresh"].(string),
		},
	)
}

func (s *service) FetchAuthUser(c *fiber.Ctx) error {
	token := c.Cookies("__token")
	if token == "" {
		log.Println("access token not found FetchAuthUser")
		return c.Status(fiber.StatusUnauthorized).JSON(presenters.AuthErrorResponse("Unauthorized"))
	}
	var session presenters.AuthSession
	if err := s.cache.Get(token, &session); err != nil {
		refresh := c.Cookies("__refresh")
		if refresh == "" {
			log.Println("refresh token not found")
			return c.Status(fiber.StatusUnauthorized).JSON(presenters.AuthErrorResponse("Unauthorized"))
		}
		newTokens, err := s.GenerateNewTokens(refresh)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusUnauthorized).JSON(presenters.AuthErrorResponse("Unauthorized"))
		}
		userSession, ok := newTokens["session"].(presenters.AuthSession)
		if !ok {
			log.Println(err)
			return c.Status(fiber.StatusUnauthorized).JSON(presenters.AuthErrorResponse("Unauthorized"))
		}
		return c.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"status":  true,
				"data":    userSession,
				"token":   newTokens["token"].(string),
				"refresh": newTokens["refresh"].(string),
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"status": true,
			"data":   session,
		},
	)
}

func (s *service) UpdatePassword(id string, request any, userType string, isOTP bool) (bool, error) {
	data := request.(*models.ChangePassword)
	switch s.userType[userType] {
	case "customer":
		user, err := s.repo.ReadCustomer(id, "id")
		if err != nil {
			return false, err
		}
		if !s.hashCheck(user.Password, data.CurrentPassword) {
			return false, fmt.Errorf("current password do not match our records")
		}
		return s.repo.UpdatePassword(user.ID, data.Password, userType, isOTP)
	case "agent":
		user, err := s.repo.ReadAgent(id, "id")
		if err != nil {
			return false, err
		}
		if !s.hashCheck(user.Password, data.CurrentPassword) {
			return false, fmt.Errorf("current password do not match our records")
		}
		return s.repo.UpdatePassword(user.ID, data.Password, userType, isOTP)
	case "merchant":
		user, err := s.repo.ReadMerchant(id, "id")
		if err != nil {
			return false, err
		}
		if !s.hashCheck(user.Password, data.CurrentPassword) {
			return false, fmt.Errorf("current password do not match our records")
		}
		return s.repo.UpdatePassword(user.ID, data.Password, userType, isOTP)

	case "asinyo":
		user, err := s.repo.ReadAdmin(id, "id")
		if err != nil {
			return false, err
		}
		if !s.hashCheck(user.Password, data.CurrentPassword) {
			return false, fmt.Errorf("current password do not match our records")
		}
		return s.repo.UpdatePassword(user.ID, data.Password, userType, isOTP)
	default:
		return false, fmt.Errorf("current password do not match our records")
	}
}

func (s *service) ResetPassword(request *models.ResetPassword, username, userType string) (bool, error) {
	switch userType {
	case "customer":
		if user, err := s.repo.ReadCustomer(username, "username"); err != nil {
			return false, err
		} else {
			return s.repo.ResetPassword(user.ID, request.NewPassword, userType)
		}
	case "agent":
		if user, err := s.repo.ReadAgent(username, "username"); err != nil {
			return false, err
		} else {
			return s.repo.ResetPassword(user.ID, request.NewPassword, userType)
		}
	case "merchant":
		if user, err := s.repo.ReadMerchant(username, "username"); err != nil {
			return false, err
		} else {
			return s.repo.ResetPassword(user.ID, request.NewPassword, userType)
		}
	case "asinyo":
		if user, err := s.repo.ReadAdmin(username, "username"); err != nil {
			return false, err
		} else {
			return s.repo.ResetPassword(user.ID, request.NewPassword, userType)
		}
	default:
		return false, fmt.Errorf("no record found")
	}
}
func (s *service) SendUserVerificationCode(username string) (string, error) {
	code, _ := application.GenerateOTP(6)
	if application.UsernameType(username, "phone") {
		s.noti.Broadcast(
			&notification.Message{
				Data: services.SMSPayload{
					Recipients: []string{username},
					Message: fmt.Sprintf(
						"\nOTP: %s\nWe appreciate your effort to join Asinyo! Enter the OTP code to proceed with your sign up.\n\nTeam Asinyo, Connecting farmers and impacting lives.\nTel: +233247770819.",
						code,
					),
				},
			},
		)
	}
	if application.UsernameType(username, "email") {
		s.noti.Broadcast(
			&notification.Message{
				Data: services.MailerMessage{
					To:       username,
					Subject:  "ASINYO SIGN UP VERIFICATION",
					Template: "verification",
					Data: struct {
						Code string
						Tel  string
					}{
						code,
						config.App().AsinyoPhone,
					},
				},
			},
		)
	}
	return code, nil
}
func (s *service) SendPasswordResetCode(username, userType string) (string, error) {
	if userType == "customer" {
		_, err := s.repo.ReadCustomer(username, "username")
		if err != nil {
			return "", err
		}
	}
	if userType == "agent" {
		_, err := s.repo.ReadAgent(username, "username")
		if err != nil {
			return "", err
		}
	}
	if userType == "merchant" {
		_, err := s.repo.ReadMerchant(username, "username")
		if err != nil {
			return "", err
		}
	}
	if userType == "asinyo" {
		_, err := s.repo.ReadAdmin(username, "username")
		if err != nil {
			return "", err
		}
	}

	code, _ := application.GenerateOTP(6)
	if application.UsernameType(username, "phone") {
		s.noti.Broadcast(
			&notification.Message{
				Data: services.SMSPayload{
					Recipients: []string{username},
					Message: fmt.Sprintf(
						"You are a step away to complete your password reset! Enter the reset code to proceed. %s",
						code,
					),
				},
			},
		)
	}
	if application.UsernameType(username, "email") {
		s.noti.Broadcast(
			&notification.Message{
				Data: services.MailerMessage{
					To:       username,
					Subject:  "ASINYO PASSWORD RESET",
					Template: "resetpassword",
					Data:     code,
				},
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

func (s *service) signinCustomer(c *fiber.Ctx, request models.User) error {
	user, err := s.repo.ReadCustomer(request.Username, "username")
	if err != nil || !s.hashCheck(user.Password, request.Password) {
		return c.Status(fiber.StatusNotFound).JSON(presenters.AuthErrorResponse("Bad credentials!"))
	}
	return c.Status(fiber.StatusOK).JSON(presenters.AuthCustomerResponse(user))
}
func (s *service) signinAgent(c *fiber.Ctx, request models.User) error {
	user, err := s.repo.ReadAgent(request.Username, "username")
	if err != nil || !s.hashCheck(user.Password, request.Password) {
		return c.Status(fiber.StatusNotFound).JSON(presenters.AuthErrorResponse("Bad credentials!"))
	}
	return c.Status(fiber.StatusOK).JSON(presenters.AuthAgentResponse(user))
}
func (s *service) signinMerchant(c *fiber.Ctx, request models.User) error {
	user, err := s.repo.ReadMerchant(request.Username, "username")
	if err != nil || !s.hashCheck(user.Password, request.Password) {
		return c.Status(fiber.StatusNotFound).JSON(presenters.AuthErrorResponse("Bad credentials!"))
	}
	return c.Status(fiber.StatusOK).JSON(presenters.AuthMerchantResponse(user))
}
func (s *service) signinAdmin(c *fiber.Ctx, request models.User) error {
	user, err := s.repo.ReadAdmin(request.Username, "username")
	if err != nil || !s.hashCheck(user.Password, request.Password) {
		return c.Status(fiber.StatusNotFound).JSON(presenters.AuthErrorResponse("Bad credentials!"))
	}
	return c.Status(fiber.StatusOK).JSON(presenters.AuthAdminResponse(user))
}

func (s *service) GenerateNewTokens(__token string) (map[string]any, error) {
	if ok := s.cache.Exist(__token); !ok {
		return nil, fmt.Errorf("refresh token not found")
	}
	var oldSession presenters.AuthSession
	if err := s.cache.Get(__token, &oldSession); err != nil {
		return nil, err
	}
	newSession, err := s.newUserSession(&oldSession)
	if err != nil {
		return nil, err
	}
	token := application.RandomString(64)
	refresh := application.RandomString(64)
	err = s.cache.Set(token, &newSession, AccessTokenExpiry)
	if err != nil {
		return nil, err
	}
	err = s.cache.Set(refresh, &newSession, RefreshTokenExpiry)
	if err != nil {
		return nil, err
	}
	// _ = s.cache.Delete(__token)
	return map[string]any{
		"token":   token,
		"refresh": refresh,
		"session": newSession,
	}, nil
}
func (s *service) newUserSession(oldSession *presenters.AuthSession) (*presenters.AuthSession, error) {

	if oldSession.UserType == "business" || oldSession.UserType == "individual" {
		user, err := s.repo.ReadCustomer(oldSession.Username, "username")
		if err != nil {
			return nil, err
		}
		return s.authCustomerResponse(user), nil
	}

	return nil, nil
}
func (s *service) authAdminResponse(data *ent.Admin) *presenters.AuthSession {
	return &presenters.AuthSession{
		ID:          data.ID,
		Username:    data.Username,
		SessionName: strings.Split(data.OtherName, " ")[0],
		Permissions: func() []string {
			roles, err := data.Edges.RolesOrErr()
			if err != nil {
				return nil
			}
			var permissions []string
			for _, role := range roles {
				perms, err := role.Edges.PermissionsOrErr()
				if err != nil {
					return nil
				}
				for _, perm := range perms {
					if lo.Contains(permissions, perm.Slug) {
						continue
					}
					permissions = append(permissions, perm.Slug)
				}
			}
			return permissions
		}(),
	}
}
func (s *service) authCustomerResponse(data *ent.Customer) *presenters.AuthSession {
	if c, err := data.Edges.IndividualOrErr(); err == nil {
		return &presenters.AuthSession{
			ID:          data.ID,
			Username:    data.Username,
			SessionName: strings.Split(c.OtherName, " ")[0],
			DisplayName: fmt.Sprintf("%s %s", c.OtherName, c.LastName),
			UserType:    data.Type,
		}
	}
	if c, err := data.Edges.BusinessOrErr(); err == nil {
		return &presenters.AuthSession{
			ID:          data.ID,
			Username:    data.Username,
			SessionName: c.Name,
			DisplayName: c.Name,
			UserType:    data.Type,
		}
	}

	return nil
}
func (s *service) authAgentResponse(data *ent.Agent) *presenters.AuthSession {
	return &presenters.AuthSession{
		ID:          data.ID,
		Username:    data.Username,
		SessionName: strings.Split(data.OtherName, " ")[0],
		DisplayName: fmt.Sprintf("%s %s", data.OtherName, data.LastName),
		UserType:    "agent",
	}
}
func (s *service) authMerchantResponse(data *ent.Merchant) *presenters.AuthSession {
	if s, err := data.Edges.SupplierOrErr(); err == nil {
		return &presenters.AuthSession{
			ID:          data.ID,
			Username:    data.Username,
			SessionName: strings.Split(s.OtherName, " ")[0],
			DisplayName: fmt.Sprintf("%s %s", s.OtherName, s.LastName),
			UserType:    data.Type,
			OTP:         data.Otp,
			Storefront:  data.Edges.Store.ID,
		}
	}
	if r, err := data.Edges.RetailerOrErr(); err == nil {
		return &presenters.AuthSession{
			ID:          data.ID,
			Username:    data.Username,
			SessionName: strings.Split(r.OtherName, " ")[0],
			DisplayName: fmt.Sprintf("%s %s", r.OtherName, r.LastName),
			UserType:    data.Type,
			OTP:         data.Otp,
			Storefront:  data.Edges.Store.ID,
		}
	}
	return nil
}
