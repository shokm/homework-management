package database

import (
	"backend/gen/models"
	"backend/handler/dotenv"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoginAuthDB(body models.AuthUserReq) error {
	// dsnをdotenvから読み込み
	dsnLoadByDotenv, err := dotenv.LoadDataSourceName()
	if err != nil {
		return errors.New("failed to read secret key")
	}
	dsn := dsnLoadByDotenv.(string)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.New("failed to connect database")
	}

	// ユーザー情報を取得
	userInfo := UserLists{}
	db.Model(&UserLists{}).Where("screen_name = ?", body.ScreenName).First(&userInfo)

	// 入力されたパスワードのハッシュ化
	b := []byte(body.Password)
	sha256 := sha256.Sum256(b)
	hashedInputPassword := hex.EncodeToString(sha256[:])

	// パスワードの比較
	if userInfo.Password != hashedInputPassword {
		sqlDB, err := db.DB()
		if err != nil {
			return errors.New(err.Error())
		}
		sqlDB.Close()
		return errors.New("wrong password")
	}

	// セッションを切る
	sqlDB, err := db.DB()
	if err != nil {
		return errors.New(err.Error())
	}
	sqlDB.Close()

	return nil
}

func CreateUserDB(body models.AuthUserReq) error {
	// dsnをdotenvから読み込み
	dsnLoadByDotenv, err := dotenv.LoadDataSourceName()
	if err != nil {
		return errors.New("failed to read secret key")
	}
	dsn := dsnLoadByDotenv.(string)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.New("failed to connect database")
	}

	// スクリーンネームが既に存在する場合、エラー
	var count int64
	db.Model(&UserLists{}).Where("screen_name = ?", body.ScreenName).Count(&count)
	if count != 0 {
		sqlDB, err := db.DB()
		if err != nil {
			return errors.New(err.Error())
		}
		sqlDB.Close()
		return errors.New("screen_name is exist")
	}

	// パスワードのハッシュ化
	b := []byte(body.Password)
	sha256 := sha256.Sum256(b)
	hashedPassword := hex.EncodeToString(sha256[:])

	// 構造体作成
	UserLists := UserLists{}

	// スクリーンネームとパスワードを代入
	UserLists.ScreenName = body.ScreenName
	UserLists.Password = hashedPassword

	// ユーザー作成
	db.Create(&UserLists)

	// セッションを切る
	sqlDB, err := db.DB()
	if err != nil {
		return errors.New(err.Error())
	}
	sqlDB.Close()

	return nil
}

func LoadUserInfoByScreenName(screenName string) (UserLists, error){
	userInfo := UserLists{}

	// dsnをdotenvから読み込み
	dsnLoadByDotenv, err := dotenv.LoadDataSourceName()
	if err != nil {
		return userInfo, errors.New("failed to read secret key")
	}
	dsn := dsnLoadByDotenv.(string)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return userInfo, errors.New(err.Error())
	}

	db.Model(&UserLists{}).Where("screen_name = ?", screenName).First(&userInfo)

	// セッションを切る
	sqlDB, err := db.DB()
	if err != nil {
		return userInfo, errors.New(err.Error())
	}
	sqlDB.Close()

	return userInfo, nil
}
