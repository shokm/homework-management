package handler

import (
	"backend/gen/models"
	"backend/gen/restapi/operations/task_api"
	"backend/handler/auth_jwt"
	"backend/handler/database"

	"github.com/go-openapi/runtime/middleware"
)

func GetSubjects(params task_api.GetSubjectsParams, principal interface{}) middleware.Responder {

	// 認証情報からユーザーIDを取得
	header := params.HTTPRequest.Header.Get("Authorization")
	returnValue, err := auth_jwt.ValidateTokenHandler(header)
	if err != nil {
		return task_api.NewGetSubjectsUnauthorized()
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
		return task_api.NewGetSubjectsServiceUnavailable()
	}

	// 結果を入れる変数
	result := models.SubjectsMultiple{}

	// 個数を入れる
	length := len(resultListsFromDB)
	result.TotalCount = int64(length)
	for i := 0; i < length; i++ {
		// tmp変数
		tmpSubjectSingle := models.SubjectSingle{}
		tmpSubjectSingle.CreatedAt = resultListsFromDB[i].CreatedAt
		tmpSubjectSingle.IsArchived = resultListsFromDB[i].IsArchived
		tmpSubjectSingle.SubjectID = int64(resultListsFromDB[i].SubjectID)
		tmpSubjectSingle.SubjectName = resultListsFromDB[i].SubjectName
		// tmp変数に入れた結果をresultに格納
		result.Subjects = append(result.Subjects, &tmpSubjectSingle)
	}

	return task_api.NewGetSubjectsOK().WithPayload(&result)
}
