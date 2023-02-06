package handler

import (
	"backend/gen/models"
	"backend/gen/restapi/operations/task_api"
	"backend/handler/auth_jwt"
	"backend/handler/database"
	"strconv"

	"github.com/go-openapi/runtime/middleware"
)

func GetTasksBySubjects(params task_api.GetTasksBySubjectsParams, principal interface{}) middleware.Responder {

	// 認証情報からユーザーIDを取得
	header := params.HTTPRequest.Header.Get("Authorization")
	returnValue, err := auth_jwt.ValidateTokenHandler(header)
	if err != nil {
		return task_api.NewGetTasksBySubjectsUnauthorized()
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
	resultListsFromDB, err := database.GetSubjectsByUserID(userID, isArchived)
	if err != nil {
		return task_api.NewGetTasksBySubjectsServiceUnavailable()
	}

	// 結果を入れる変数
	result := models.ReturnTasksBySubjects{}

	// 個数を入れる
	length := len(resultListsFromDB)
	result.TotalCount = int64(length)
	for i := 0; i < length; i++ {
		// tmp変数
		tmpSubjectSingle := models.ReturnTasksBySubject{}
		tmpSubjectSingle.CreatedAt = resultListsFromDB[i].CreatedAt
		tmpSubjectSingle.IsArchived = resultListsFromDB[i].IsArchived
		tmpSubjectSingle.SubjectID = int64(resultListsFromDB[i].SubjectID)
		tmpSubjectSingle.SubjectName = resultListsFromDB[i].SubjectName

		// 課題の取得
		resultTasksFromDB, err := database.GetTasksBySubjectID(userID, tmpSubjectSingle.SubjectID, isArchived)
		if err != nil {
			return task_api.NewGetTasksBySubjectsServiceUnavailable()
		}
		// 結果を詰めていく
		lengthTask := len(resultTasksFromDB)
		tmpSubjectSingle.TotalCount = int64(lengthTask)
		for j := 0; j < lengthTask; j++ {
			// tmp変数
			tmpTaskSingle := models.TaskSingle{}
			tmpTaskSingle.CreatedAt = strconv.FormatInt(resultTasksFromDB[j].CreatedAt.Unix(), 10)
			tmpTaskSingle.DeadlineAt = strconv.FormatInt(resultTasksFromDB[j].DeadlineAt.Unix(), 10)
			tmpTaskSingle.IsArchived = resultTasksFromDB[j].IsArchived
			tmpTaskSingle.SubjectID = int64(resultTasksFromDB[j].SubjectID)
			tmpTaskSingle.SubjectName = tmpSubjectSingle.SubjectName
			tmpTaskSingle.StateID = int64(resultTasksFromDB[j].StateID)
			tmpTaskSingle.TaskDescription = resultTasksFromDB[j].TaskDescription
			tmpTaskSingle.TaskID = int64(resultTasksFromDB[j].TaskID)
			tmpTaskSingle.TaskName = resultTasksFromDB[j].TaskName
			// tmp変数に入れた結果をresultに格納
			tmpSubjectSingle.Tasks = append(tmpSubjectSingle.Tasks, &tmpTaskSingle)
		}

		// tmp変数に入れた結果をresultに格納
		result.Subjects = append(result.Subjects, &tmpSubjectSingle)
	}

	return task_api.NewGetTasksBySubjectsOK().WithPayload(&result)
}
