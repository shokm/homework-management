package auth

import (
	"backend/gen/models"
	"backend/handler/database"
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GetTokenHandler(body database.UserLists) models.AuthReturnJWT {

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

	// トークンに署名を付与
	tokenString, err := token.SignedString([]byte("SECRET_KEY"))
	if err != nil {}

	result.Token = tokenString

	return result
}

func ValidateTokenHandler(tokenHeader string) (interface{}, error) {

	// ヘッダからbearer...を切り離す
	tokenString := strings.Split(tokenHeader, " ")[1]

	// 認証
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		// TODO: キーを設定
		return []byte("SECRET_KEY"), nil
	})

	result := models.AuthReturnUser{}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		result.UserID = int64(claims["user_id"].(float64))
		result.ScreenName = string(claims["screen_name"].(string))
	} else {
		// TODO: エラーを正しく書く
		return nil, errors.New(err.Error())
	}

	return result, nil
}
