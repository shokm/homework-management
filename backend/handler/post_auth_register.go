package handler

import (
	"backend/gen/restapi/operations/auth_api"
	"backend/handler/auth_jwt"
	"backend/handler/database"

	"github.com/go-openapi/runtime/middleware"
)

func PostAuthRegister(params auth_api.PostAuthRegisterParams) middleware.Responder {

	// TODO: バリデーションをかける
	// スクリーンネームとパスワードが空だった場合、エラー
	if params.Body.ScreenName == "" {
		return auth_api.NewPostAuthLoginUnauthorized()
	}
	if params.Body.Password == "" {
		return auth_api.NewPostAuthLoginUnauthorized()
	}

	// ユーザー作成
	err := database.CreateUserDB(*params.Body)
	if err != nil {
		return auth_api.NewPostAuthRegisterUnauthorized()
	}

	// DBからユーザー情報取得
	userInfo, err := database.LoadUserInfoByScreenName(params.Body.ScreenName)
	if err != nil {
		return auth_api.NewPostAuthRegisterInternalServerError()
	}

	// トークン取得
	result, err := auth_jwt.GetTokenHandler(userInfo)
	if err != nil {
		return auth_api.NewPostAuthRegisterUnauthorized()
	}

	return auth_api.NewPostAuthRegisterOK().WithPayload(&result)
}
