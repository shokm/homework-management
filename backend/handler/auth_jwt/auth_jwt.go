package auth_jwt

import (
	"backend/gen/models"
	"backend/handler/database"
	"backend/handler/dotenv"
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GetTokenHandler(body database.UserLists) (models.AuthReturnJWT, error) {

	// 返却値
	result := models.AuthReturnJWT{}

	// ペイロードの作成
	claims := jwt.MapClaims{
		"user_id": body.UserID,
		"screen_name": body.ScreenName,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	// トークン生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// キーの読み込み
	seacretKeyLoadByDotenv, err := dotenv.LoadJWTSecretKey()
	if err != nil {
		return result, errors.New("failed to read secret key")
	}
	// トークンに署名を付与
	tokenString, err := token.SignedString([]byte(seacretKeyLoadByDotenv.(string)))
	if err != nil {
		return result, errors.New("signing failure")
	}

	result.Token = tokenString

	return result, nil
}

func ValidateTokenHandler(tokenHeader string) (interface{}, error) {

	// ヘッダからbearer...を切り離す
	tokenString := strings.Split(tokenHeader, " ")[1]

	// 認証
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		// キーの読み込み
		seacretKeyLoadByDotenv, err := dotenv.LoadJWTSecretKey()
		if err != nil {
			return nil, errors.New("failed to read secret key")
		}
		// 署名を設定
		return []byte(seacretKeyLoadByDotenv.(string)), nil
	})

	result := models.AuthReturnUser{}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		result.UserID = int64(claims["user_id"].(float64))
		result.ScreenName = string(claims["screen_name"].(string))
	} else {
		return nil, errors.New(err.Error())
	}

	return result, nil
}
