package auth

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"log"
	"strings"
	"time"
)

type JwtClaim struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	jwt.RegisteredClaims
}

func GenerateToken(email, firstName, lastName string) (string, error) {
	expirationTime := jwt.NewNumericDate(time.Now().Add(1 * time.Hour))
	claims := &JwtClaim{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expirationTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(viper.GetString("auth.jwtKey")))
	return tokenString, err
}

func CheckAuth(ctx *fiber.Ctx) error {
	headers := ctx.GetReqHeaders()
	token := headers["Authorization"]
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	token = strings.Split(token, "Bearer ")[1]
	err := validateToken(token)
	log.Println(err)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Next()
}

func validateToken(signedToken string) error {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("auth.jwtKey")), nil
	})
	if err != nil {
		return err
	}

	fmt.Println(token.Claims.(jwt.MapClaims))

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token or expired")
}
