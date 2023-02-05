package handler

import (
	"backend/gen/models"
	"backend/gen/restapi/operations/task_api"
	"backend/handler/auth_jwt"
	"backend/handler/database"
	"strconv"

	"github.com/go-openapi/runtime/middleware"
)

func GetSubjectBySubjectID(params task_api.GetSubjectBySubjectIDParams, principal interface{}) middleware.Responder {

	// 認証情報からユーザーIDを取得
	header := params.HTTPRequest.Header.Get("Authorization")
	returnValue, err := auth_jwt.ValidateTokenHandler(header)
	if err != nil {
		return task_api.NewGetSubjectBySubjectIDUnauthorized()
	}
	userInfo := returnValue.(models.AuthReturnUser)
	userID := userInfo.UserID

	// DBから結果を取得
	resultListFromDB, err := database.GetSubjectBySubjectID(userID, params.SubjectID)
	if err != nil {
		return task_api.NewGetSubjectBySubjectIDServiceUnavailable()
	}

	// 結果を入れる変数
	result := models.SubjectSingle{}

	// 結果を詰めていく
	result.CreatedAt = strconv.FormatInt(resultListFromDB.CreatedAt.Unix(), 10)
	result.IsArchived = resultListFromDB.IsArchived
	result.SubjectID = int64(resultListFromDB.SubjectID)
	result.SubjectName = resultListFromDB.SubjectName

	return task_api.NewGetSubjectBySubjectIDOK().WithPayload(&result)
}
