package handler

import (
	"backend/gen/models"
	"backend/gen/restapi/operations/task_api"
	"backend/handler/auth_jwt"
	"backend/handler/database"
	"strconv"

	"github.com/go-openapi/runtime/middleware"
)

func GetTaskByTaskID(params task_api.GetTaskByTaskIDParams, principal interface{}) middleware.Responder {

	// 認証情報からユーザーIDを取得
	header := params.HTTPRequest.Header.Get("Authorization")
	returnValue, err := auth_jwt.ValidateTokenHandler(header)
	if err != nil {
		return task_api.NewGetTaskByTaskIDUnauthorized()
	}
	userInfo := returnValue.(models.AuthReturnUser)
	userID := userInfo.UserID

	// リクエストされたtaskIDが該当ユーザーのものか・taskIDが存在するかどうか検索
	resultTaskIDCount, err := database.GetCountTaskByTaskID(userID, int64(params.TaskID))
	if err != nil {
		return task_api.NewGetTaskByTaskIDInternalServerError()
	}
	if resultTaskIDCount == 0 {
		// IDが存在しない場合
		return task_api.NewGetTaskByTaskIDNotFound()
	}

	// DBから結果を取得
	resultListFromDB, err := database.GetTaskByTaskID(userID, params.TaskID)
	if err != nil {
		return task_api.NewGetTaskByTaskIDServiceUnavailable()
	}

	// 結果を入れる変数
	result := models.TaskSingle{}

	// 結果を詰めていく
	result.CreatedAt = strconv.FormatInt(resultListFromDB.CreatedAt.Unix(), 10)
	result.DeadlineAt = strconv.FormatInt(resultListFromDB.DeadlineAt.Unix(), 10)
	result.IsArchived = resultListFromDB.IsArchived
	result.SubjectID = int64(resultListFromDB.SubjectID)
	result.SubjectName = resultListFromDB.SubjectName
	result.StateID = int64(resultListFromDB.StateID)
	result.TaskDescription = resultListFromDB.TaskDescription
	result.TaskID = int64(resultListFromDB.TaskID)
	result.TaskName = resultListFromDB.TaskName

	return task_api.NewGetTaskByTaskIDOK().WithPayload(&result)
}
