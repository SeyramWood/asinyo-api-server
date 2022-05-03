package auth

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/config"
	"github.com/SeyramWood/ent"
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
	var request models.Auth
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presenters.AuthErrorResponse(err))
	}
	if er := validate(&request); er != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
	}

	if user, err := s.fetchUser(request.Username); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(presenters.AuthErrorResponse("No user found"))
	} else {
		if s.hashCheck(user.Password, request.Password) {
			claim := jwt.NewWithClaims(jwt.SigningMethodHS256, s.claims(user.ID))

			token, err := claim.SignedString([]byte(config.App().Key))

			if err != nil {
				return c.SendStatus(fiber.StatusInternalServerError)
			}
			c.Cookie(&fiber.Cookie{
				Name:     "token",
				Value:    token,
				Expires:  time.Now().Add(time.Hour * 24 * 7),
				HTTPOnly: true,
			})
			return c.Status(fiber.StatusOK).JSON(presenters.AuthSuccessResponse(user))
		}

	}
	return c.Status(fiber.StatusUnauthorized).JSON(presenters.AuthErrorResponse("Bad credentials"))
}

func (s *service) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}

func (s *service) fetchUser(username string) (*ent.User, error) {
	return s.repo.Read(username)
}
func (s *service) hashCheck(hash []byte, plain string) bool {
	if err := bcrypt.CompareHashAndPassword(hash, []byte(plain)); err != nil {
		return false
	}
	return true
}

func (s *service) claims(id int) jwt.RegisteredClaims {
	return jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(id),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
	}
}

func validate(i interface{}) map[string]string {
	typ := reflect.TypeOf(i)
	value := reflect.ValueOf(i)
	var err = make(map[string]string)
	var structure = make(map[string]string)
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	for i := 0; i < typ.NumField(); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			field := strings.ToLower(typ.Field(i).Name)
			msg := validateByTag(field, typ.Field(i).Type.String(), value.Field(i), typ.Field(i).Tag.Get("validate"))
			if len(msg) > 0 {
				mut.Lock()
				err[field] = msg
				mut.Unlock()
			}
			mut.Lock()
			structure[field] = msg
			mut.Unlock()
		}(i)
	}
	wg.Wait()
	if len(err) > 0 {
		return structure
	}
	return nil
}
func validateByTag(field string, fieldType string, v reflect.Value, tag string) string {

	var err string
	switch fieldType {
	case "string":
		checkEmail := false
		for _, v := range strings.Split(tag, "|") {
			if v == "email" {
				checkEmail = true
				break
			}
		}
		if checkEmail {
			break
		}
		r, _ := regexp.Compile("^[0-9a-zA-Z ]+$")
		if !r.MatchString(v.String()) {
			err = fmt.Sprintf("The %s must be a string.", field)
		}
	}
	for _, rule := range strings.Split(tag, "|") {
		if strings.Contains(rule, ":") {
			r := strings.Split(rule, ":")
			switch string(r[0]) {
			case "max":
				value, _ := strconv.Atoi(r[1])
				if len(v.String()) > value {
					err = fmt.Sprintf("The %s must not be greater than %v characters", field, value)

				}
			case "min":
				value, _ := strconv.Atoi(r[1])
				if len(v.String()) < value {
					err = fmt.Sprintf("The %s must be at least %v characters", field, value)
				}
			}
			continue
		}
		switch rule {
		case "required":
			if !v.IsValid() {
				err = fmt.Sprintf("The %s field is required", field)
			}
		case "email":
			r, _ := regexp.Compile("^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$")
			if !r.MatchString(v.String()) {
				err = fmt.Sprintf("The %s must be a valid email address.", field)
			}
		}
	}
	return err
}
