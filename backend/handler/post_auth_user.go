package handler

import (
	"backend/gen/models"
	"backend/gen/restapi/operations/auth_api"
	"backend/handler/auth_jwt"

	"github.com/go-openapi/runtime/middleware"
)

func PostAuthUser(params auth_api.PostAuthUserParams, principal interface{}) middleware.Responder {

	header := params.HTTPRequest.Header.Get("Authorization")

	// TODO: エラーハンドリング
	returnValue, err := auth_jwt.ValidateTokenHandler(header)
	if err != nil {
		return auth_api.NewPostAuthUserUnauthorized()
	}

	userInfo := returnValue.(models.AuthReturnUser)

	return auth_api.NewPostAuthUserOK().WithPayload(&userInfo)
}
