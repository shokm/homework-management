package handler

import (
	"backend/gen/restapi/operations/auth_api"
	"backend/handler/auth_jwt"
	"backend/handler/database"

	"github.com/go-openapi/runtime/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostAuthLogin(params auth_api.PostAuthLoginParams) middleware.Responder {

	// ユーザー認証
	err := database.LoginAuthDB(*params.Body)
	if err != nil {
		// TODO: middleware.NotImplementedを置き換える
		return auth_api.NewPostAuthLoginUnauthorized()
	}

	// ユーザー情報取得
	dsn := "host=localhost user=postgres password=postgrespw dbname=postgres port=55000 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	userInfo := database.UserLists{}
	db.Model(&database.UserLists{}).Where("screen_name = ?", params.Body.ScreenName).First(&userInfo)

	// セッションを切る
	sqlDB, err := db.DB()
	if err != nil {
		return auth_api.NewPostAuthLoginInternalServerError()
	}
	sqlDB.Close()

	// トークン取得
	result, err := auth_jwt.GetTokenHandler(userInfo)
	if err != nil {
		return auth_api.NewPostAuthLoginUnauthorized()
	}

	// return operations.NewHelloOK().WithPayload("Hello Swagger")
	return auth_api.NewPostAuthLoginOK().WithPayload(&result)
}
