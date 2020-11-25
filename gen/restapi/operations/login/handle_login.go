package login

import (
	loginUser "medicarePartD/bus/login"
	"medicarePartD/bus/responder"
	"medicarePartD/gen/models"
	"strings"

	"github.com/go-openapi/runtime/middleware"
)

func HandleLogin(params LoginParams) middleware.Responder {
	if params.Login != nil {
		err := loginUser.LoginUser(params.Login.Email, params.Login.Password)

		if err != nil {

			if strings.Contains(err.Error(), "unauthorized") {
				return NewLoginUnauthorized().WithPayload(&models.ResponseCode{
					Code:    "401",
					Message: "you do not have the proper credntials to access this database"})
			}
			if strings.Contains(err.Error(), "forbidden") {
				return NewLoginForbidden().WithPayload(&models.ResponseCode{
					Code:    "403",
					Message: "incorrect username or password"})
			}
			if strings.Contains(err.Error(), "not acceptable") {
				return NewLoginNotAcceptable().WithPayload(&models.ResponseCode{
					Code:    "406",
					Message: "please enter a valid email"})
			}
			return NewLoginInternalServerError().WithPayload(&models.ResponseCode{
				Code:    "500",
				Message: "server error"})

		}
		return responder.CookieHandler((NewLoginOK().WithPayload(
			&models.JwtCode{
				Code:    "200",
				Jwt:     responder.CreateToken(params.Login.Email)[2],
				Message: "successful Login"})), params.Login.Email)
	}
	return NewLoginInternalServerError().WithPayload(&models.ResponseCode{
		Code:    "500",
		Message: "server error"})
}
