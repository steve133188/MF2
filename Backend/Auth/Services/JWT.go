package Services

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"os"
	"time"
)

var mySigningKey = []byte(os.Getenv("SECRET_KEY"))

func GenJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "TCO"
	//claims["aud"] = "billing.jwtgo.io"
	claims["iss"] = "matrixsense.io"
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func AuthCheck(c * fiber.Ctx) error{
	validToken, err := GenJWT()
	fmt.Println(validToken)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Please Login", // invalid token
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message":"authenticated",
		"data": fiber.Map{
			"token":string(validToken),  // token
		},
	})
}

