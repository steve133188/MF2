package Middleware

import (
	"aws-lambda-chatroom/Handler"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/dgrijalva/jwt-go"
)

var token_pwd string = "51c3d3fc-3e15-4c19-7437-d74f5e5f906c"

type EventHandler func(events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error)

func JWTHandler(next EventHandler) MiddlwareHandler {
	return func(req events.APIGatewayProxyRequest) (interface{}, error) {
		token := req.Headers["authorization"]
		var errMessage string

		if len(token) != 0 {
			token = strings.Split(token, "Bearer ")[1]

			jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
				Secret := token_pwd
				return []byte(Secret), nil
			})
			if err == nil && jwtToken != nil {
				if _, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
					return next(req)
				}
			}

			if err != nil && jwt.ValidationErrorExpired != 0 {
				errMessage = "JWT Token is Expired"
			} else if err != nil && jwt.ValidationErrorSignatureInvalid != 0 {
				errMessage = "Invalid JWT Token"
			} else {
				errMessage = "Cannot Handle This Token"
			}

		} else {
			errMessage = "Missing JWT Token"
		}

		return Handler.ApiResponse(http.StatusUnauthorized, map[string]string{
			"message": errMessage,
		}), nil
	}
}
