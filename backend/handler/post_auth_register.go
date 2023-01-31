package handler

import (
	"backend/gen/restapi/operations/auth_api"
	"backend/handler/auth_jwt"
	"backend/handler/database"

	"github.com/go-openapi/runtime/middleware"
)

func PostAuthRegister(params auth_api.PostAuthRegisterParams) middleware.Responder {

	// ユーザー作成
	err := database.CreateUserDB(*params.Body)
	if err != nil {
		return auth_api.NewPostAuthRegisterUnauthorized()
	}

	// DBからユーザー情報取得
	userInfo, err := database.LoadUserInfoByScreenName(params.Body.ScreenName)
	if err != nil {
		return auth_api.NewPostAuthUserInternalServerError()
	}

	// トークン取得
	result, err := auth_jwt.GetTokenHandler(userInfo)
	if err != nil {
		return auth_api.NewPostAuthRegisterUnauthorized()
	}

	return auth_api.NewPostAuthRegisterOK().WithPayload(&result)
}
