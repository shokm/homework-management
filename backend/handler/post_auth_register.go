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

func CreateUserDB(body models.AuthUserReq) error {
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

	// スクリーンネームが既に存在する場合、エラー
	var count int64
	db.Model(&database.UserLists{}).Where("screen_name = ?", body.ScreenName).Count(&count)
	if count != 0 {
		sqlDB, err := db.DB()
		if err != nil {}
		sqlDB.Close()
		return errors.New("screen_name is exist")
	}

	// パスワードのハッシュ化
	b := []byte(body.Password)
	sha256 := sha256.Sum256(b)
    hashedPassword := hex.EncodeToString(sha256[:])

	// 構造体作成
	UserLists := database.UserLists{}

	// スクリーンネームとパスワードを代入
	UserLists.ScreenName = body.ScreenName
	UserLists.Password = hashedPassword

	// ユーザー作成
	db.Create(&UserLists)

	// セッションを切る
	sqlDB, err := db.DB()
	if err != nil {}
	sqlDB.Close()

	return nil
}

func PostAuthRegister(params auth_api.PostAuthRegisterParams) middleware.Responder {

	// ユーザー作成
	err := CreateUserDB(*params.Body)
	if err != nil {
		// TODO: middleware.NotImplementedを置き換える
		return middleware.NotImplemented(err.Error())
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

	return auth_api.NewPostAuthRegisterOK().WithPayload(&result)
}
