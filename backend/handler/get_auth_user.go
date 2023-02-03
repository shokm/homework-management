package handler

import (
	"backend/gen/models"
	"backend/gen/restapi/operations/auth_api"
	"backend/handler/auth_jwt"

	"github.com/go-openapi/runtime/middleware"
)

func GetAuthUser(params auth_api.GetAuthUserParams, principal interface{}) middleware.Responder {

	header := params.HTTPRequest.Header.Get("Authorization")

	returnValue, err := auth_jwt.ValidateTokenHandler(header)
	if err != nil {
		return auth_api.NewGetAuthUserUnauthorized()
	}

	userInfo := returnValue.(models.AuthReturnUser)

	return auth_api.NewGetAuthUserOK().WithPayload(&userInfo)
}
