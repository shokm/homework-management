package handler

import (
	"backend/gen/models"
	"backend/gen/restapi/operations/auth_api"
	"backend/handler/auth"
	"backend/handler/database"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/go-openapi/runtime/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoginAuthDB(body models.AuthUserReq) error {
	dsn := "host=localhost user=postgres password=postgrespw dbname=postgres port=55000 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// スクリーンネームとパスワードが空だった場合、エラー
	if body.ScreenName == "" {
		sqlDB, err := db.DB()
		if err != nil {}
		sqlDB.Close()
		return errors.New("screen_name is empty")
	}
	if body.Password == "" {
		sqlDB, err := db.DB()
		if err != nil {}
		sqlDB.Close()
		return errors.New("password is empty")
	}

	// ユーザー情報を取得
	userInfo := database.UserLists{}
	db.Model(&database.UserLists{}).Where("screen_name = ?", body.ScreenName).First(&userInfo)

	// 入力されたパスワードのハッシュ化
	b := []byte(body.Password)
	sha256 := sha256.Sum256(b)
    hashedInputPassword := hex.EncodeToString(sha256[:])

	// パスワードの比較
	if userInfo.Password != hashedInputPassword {
		sqlDB, err := db.DB()
		if err != nil {}
		sqlDB.Close()
		return errors.New("wrong password")
	}

	// セッションを切る
	sqlDB, err := db.DB()
	if err != nil {}
	sqlDB.Close()

	return nil
}

func PostAuthLogin(params auth_api.PostAuthLoginParams) middleware.Responder {

	// ユーザー認証
	err2 := LoginAuthDB(*params.Body)
	if err2 != nil {
		// TODO: middleware.NotImplementedを置き換える
		return middleware.NotImplemented(err2.Error())
	}

	dsn := "host=localhost user=postgres password=postgrespw dbname=postgres port=55000 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	userInfo := database.UserLists{}
	db.Model(&database.UserLists{}).Where("screen_name = ?", params.Body.ScreenName).First(&userInfo)

	// トークン取得
	result := auth.GetTokenHandler(userInfo)

	// return operations.NewHelloOK().WithPayload("Hello Swagger")
	return auth_api.NewPostAuthLoginOK().WithPayload(&result)
}
