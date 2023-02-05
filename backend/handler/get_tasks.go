package handler

import (
	"backend/gen/models"
	"backend/gen/restapi/operations/task_api"
	"backend/handler/auth_jwt"
	"backend/handler/database"
	"strconv"

	"github.com/go-openapi/runtime/middleware"
)

func GetTasks(params task_api.GetTasksParams, principal interface{}) middleware.Responder {

	// 認証情報からユーザーIDを取得
	header := params.HTTPRequest.Header.Get("Authorization")
	returnValue, err := auth_jwt.ValidateTokenHandler(header)
	if err != nil {
		return task_api.NewGetTasksUnauthorized()
	}
	userInfo := returnValue.(models.AuthReturnUser)
	userID := userInfo.UserID

	// isArchivedパラメータによる条件分岐
	var isArchived (bool)
	if (params.IsArchived == nil){
		// IsArchivedパラメータがない場合、有効な課題のみ（isArchived = false）
		isArchived = false
	} else {
		// IsArchivedパラメータが存在する場合
		paramsIsArchived := *params.IsArchived
		if (paramsIsArchived) {
			// paramsIsArchivedパラメータがtrueの場合、アーカイブ済みの課題のみ（isArchived = true）
			isArchived = true
		} else {
			// paramsIsArchivedパラメータがfalseの場合、有効な課題のみ（isArchived = true）
			isArchived = false
		}
	}

	// DBから結果を取得
	resultListsFromDB, err := database.GetTasksByUserID(userID, isArchived)
	if err != nil {
		return task_api.NewGetTasksServiceUnavailable()
	}

	// 結果を入れる変数
	result := models.TasksMultiple{}

	// 結果を詰めていく
	length := len(resultListsFromDB)
	result.TotalCount = int64(length)
	for i := 0; i < length; i++ {
		// tmp変数
		tmpTaskSingle := models.TaskSingle{}
		tmpTaskSingle.CreatedAt = strconv.FormatInt(resultListsFromDB[i].CreatedAt.Unix(), 10)
		tmpTaskSingle.DeadlineAt = strconv.FormatInt(resultListsFromDB[i].DeadlineAt.Unix(), 10)
		tmpTaskSingle.IsArchived = resultListsFromDB[i].IsArchived
		tmpTaskSingle.SubjectID = int64(resultListsFromDB[i].SubjectID)
		tmpTaskSingle.SubjectName = resultListsFromDB[i].SubjectName
		tmpTaskSingle.StateID = int64(resultListsFromDB[i].StateID)
		tmpTaskSingle.TaskDescription = resultListsFromDB[i].TaskDescription
		tmpTaskSingle.TaskID = int64(resultListsFromDB[i].TaskID)
		tmpTaskSingle.TaskName = resultListsFromDB[i].TaskName
		// tmp変数に入れた結果をresultに格納
		result.Tasks = append(result.Tasks, &tmpTaskSingle)
	}

	return task_api.NewGetTasksOK().WithPayload(&result)
}
