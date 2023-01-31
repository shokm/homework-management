package dotenv

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func LoadJWTSecretKey() (interface{}, error) {

	// .envのパス指定
	err := godotenv.Load("settings/.env")
	if err != nil {
		return nil, errors.New("failed to read env")
	}

	JWTSecretKey := os.Getenv("JWT_SECRET_KEY")

	return JWTSecretKey, nil
}

func LoadDataSourceName() (interface{}, error) {

	// .envのパス指定
	err := godotenv.Load("settings/.env")
	if err != nil {
		return nil, errors.New("failed to read env")
	}

	dataSourceName := os.Getenv("DATA_SOURCE_NAME")

	return dataSourceName, nil
}
