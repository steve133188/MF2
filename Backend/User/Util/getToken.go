package Util

import (
	"github.com/dgrijalva/jwt-go"
)

func ParseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		Secret := GoDotEnvVariable("Token_pwd")
		return []byte(Secret), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}

// func AuthRequired() func(c *fiber.Ctx) {
// 	return jwtware.New(jwtware.Config{
// 		ErrorHandler: func(c *fiber.Ctx, err error) {
// 			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"error": "Unauthorized",
// 			})
// 		},
// 		SigningKey: []byte(GoDotEnvVariable("Token_pwd")),
// 	})
// }
