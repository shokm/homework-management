package handler

import (
	"backend/gen/restapi/operations/auth_api"
	"backend/handler/auth_jwt"
	"backend/handler/database"

	"github.com/go-openapi/runtime/middleware"
)

func PostAuthLogin(params auth_api.PostAuthLoginParams) middleware.Responder {

	// TODO: バリデーションをかける
	// スクリーンネームとパスワードが空だった場合、エラー
	if params.Body.ScreenName == "" {
		return auth_api.NewPostAuthLoginUnauthorized()
	}
	if params.Body.Password == "" {
		return auth_api.NewPostAuthLoginUnauthorized()
	}

	// ユーザー認証
	err := database.LoginAuthDB(*params.Body)
	if err != nil {
		// TODO: middleware.NotImplementedを置き換える
		return auth_api.NewPostAuthLoginUnauthorized()
	}

	// DBからユーザー情報取得
	userInfo, err := database.LoadUserInfoByScreenName(params.Body.ScreenName)
	if err != nil {
		return auth_api.NewPostAuthLoginInternalServerError()
	}

	// トークン取得
	result, err := auth_jwt.GetTokenHandler(userInfo)
	if err != nil {
		return auth_api.NewPostAuthLoginUnauthorized()
	}

	// return operations.NewHelloOK().WithPayload("Hello Swagger")
	return auth_api.NewPostAuthLoginOK().WithPayload(&result)
}
