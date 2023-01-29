package auth

import (
	"backend/gen/models"
	"backend/handler/database"
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
