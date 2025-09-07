package helper

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go-ecommerce-app/internal/domain"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type Auth struct {
	Secret string
}

func SetupAuth(s string) Auth {
	return Auth{
		Secret: s,
	}
}

func (a Auth) CreateHashedPassword(password string) (string, error) {
	if len(password) < 8 {
		return "", errors.New("password must be at least 8 characters")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("error hashing password")
	}
	return string(hashedPassword), nil
}

func (a Auth) GenerateToken(id uint, email string, role string) (string, error) {
	if id == 0 || email == "" || role == "" {
		return "", errors.New("id and email and role are required to generate token")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(a.Secret))
	if err != nil {
		return "", errors.New("error signing token")
	}

	return tokenString, nil
}

func (a Auth) VerifyPassword(password string, hashedPassword string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.New("password does not match")
	}

	return nil
}

func (a Auth) VerifyToken(t string) (domain.User, error) {
	tokenArr := strings.Split(t, " ")
	if len(tokenArr) != 2 {
		return domain.User{}, errors.New("invalid token")
	}

	if tokenArr[0] != "Bearer" {
		return domain.User{}, errors.New("invalid token")
	}

	tokenString := tokenArr[1]

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(a.Secret), nil
	})

	if err != nil {
		return domain.User{}, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return domain.User{}, errors.New("token is expired")
		}

		user := domain.User{
			ID:       uint(claims["user_id"].(float64)),
			Email:    claims["email"].(string),
			UserType: claims["role"].(string),
		}

		return user, nil
	}

	return domain.User{}, nil
}

func (a Auth) Authorize(ctx *fiber.Ctx) error {
	// Get the Authorization header as a string
	auth := ctx.Get("Authorization") // same as ctx.GetReqHeaders()["Authorization"] but nicer

	if strings.TrimSpace(auth) == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"message": "unauthorized",
			"reason":  "missing Authorization header",
		})
	}

	user, err := a.VerifyToken(auth)
	if err != nil || user.ID <= 0 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"message": "unauthorized",
			"reason": func() string {
				if err != nil {
					return err.Error()
				}
				return "invalid user"
			}(),
		})
	}

	ctx.Locals("user", user)
	return ctx.Next()
}

func (a Auth) GetCurrentUser(ctx *fiber.Ctx) domain.User {
	user := ctx.Locals("user")

	return user.(domain.User)
}

func (a Auth) GenerateCode() (int, error) {
	return RandomNumbers(6)
	return 0, nil
}
