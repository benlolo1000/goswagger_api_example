package register

import (
	registerUser "medicarePartD/bus/register"
	"medicarePartD/gen/models"
	"strings"

	"github.com/go-openapi/runtime/middleware"
)

func HandleRegister(params RegisterParams) middleware.Responder {
	if params.Register != nil {
		err := registerUser.RegisterUser(params.Register.Email, params.Register.Password, params.Register.ConfirmPassword)

		if err != nil {

			if strings.Contains(err.Error(), "unauthorized") {
				return NewRegisterUnauthorized().WithPayload(&models.ResponseCode{
					Code:    "401",
					Message: "unauthorized"})
			}
			if strings.Contains(err.Error(), "not acceptable") {
				return NewRegisterNotAcceptable().WithPayload(&models.ResponseCode{
					Code:    "406",
					Message: "please enter a valid email and make sure your passwords match"})
			}
			if strings.Contains(err.Error(), "conflict") {
				return NewRegisterConflict().WithPayload(&models.ResponseCode{
					Code:    "409",
					Message: "this email is already registered"})
			}
			return NewRegisterInternalServerError().WithPayload(&models.ResponseCode{
				Code:    "500",
				Message: "server error"})
		}
		return NewRegisterOK().WithPayload(&models.ResponseCode{
			Code:    "200",
			Message: "successful Registration"})
	}
	return NewRegisterInternalServerError().WithPayload(&models.ResponseCode{
		Code:    "500",
		Message: "server error"})
}
