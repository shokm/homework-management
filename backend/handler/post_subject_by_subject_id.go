package handler

import (
	"backend/gen/models"
	"backend/gen/restapi/operations/task_api"
	"backend/handler/auth_jwt"
	"backend/handler/database"
	"errors"
	"strconv"

	"github.com/go-openapi/runtime/middleware"
)

func PostSubjectBySubjectID(params task_api.PostSubjectBySubjectIDParams, principal interface{}) middleware.Responder {

	// 認証情報からユーザーIDを取得
	header := params.HTTPRequest.Header.Get("Authorization")
	returnValue, err := auth_jwt.ValidateTokenHandler(header)
	if err != nil {
		return task_api.NewPostSubjectBySubjectIDUnauthorized()
	}
	userInfo := returnValue.(models.AuthReturnUser)
	userID := userInfo.UserID

	// DBからの結果を入れる構造体
	resultListFromDB := database.SubjectLists{}

	// 結果を入れる構造体
	result := models.SubjectSingle{}

	var errorValue (error)

	// 更新と追加のAPIを共有しているため、条件分岐を行う
	resultTaskIDCount, err := database.GetCountSubjectBySubjectID(userID, int64(params.SubjectID))
	if err != nil {
		return task_api.NewPostSubjectBySubjectIDInternalServerError()
	}
	if resultTaskIDCount == 0 {
		// IDが存在しない場合
		resultListFromDB, errorValue = addSubjectValue(*params.Body, userID)
	} else {
		// IDが存在する場合
		resultListFromDB, errorValue = updSubjectValue(*params.Body, userID, int64(params.SubjectID))
	}

	if errorValue != nil {
		return task_api.NewPostSubjectBySubjectIDInternalServerError()
	}

	// 結果を詰めていく
	result.CreatedAt = strconv.FormatInt(resultListFromDB.CreatedAt.Unix(), 10)
	result.IsArchived = resultListFromDB.IsArchived
	result.SubjectID = int64(resultListFromDB.SubjectID)
	result.SubjectName = resultListFromDB.SubjectName

	return task_api.NewPostSubjectBySubjectIDOK().WithPayload(&result)
}

// 値を追加する場合
func addSubjectValue(postData models.SubjectSingle, userID int64) (database.SubjectLists, error) {

	// 書き込む値を入れる変数
	postSubjectList := database.SubjectLists{}

	postSubjectList.IsArchived = postData.IsArchived
	postSubjectList.SubjectName = postData.SubjectName
	postSubjectList.SubjectID = uint(postData.SubjectID)
	postSubjectList.UserID = uint(userID)

	// DBに書き込み
	resultSubjectID, err := database.PostSubjectBySubjectID(postSubjectList)
	if err != nil {
		return database.SubjectLists{}, errors.New(err.Error())
	}

	// DBから結果を取得
	resultListFromDB, err := database.GetSubjectBySubjectID(userID, resultSubjectID)
	if err != nil {
		return database.SubjectLists{}, errors.New(err.Error())
	}

	return resultListFromDB, nil
}

// 値を更新する場合
func updSubjectValue(postData models.SubjectSingle, userID int64, subjectID int64) (database.SubjectLists, error) {

	// 書き込む値を入れる変数
	postSubjectList := database.SubjectLists{}

	postSubjectList.IsArchived = postData.IsArchived
	postSubjectList.SubjectName = postData.SubjectName
	postSubjectList.SubjectID = uint(postData.SubjectID)

	// DBを更新
	resultSubjectID, err := database.PutSubjectBySubjectID(postSubjectList, userID, subjectID)
	if err != nil {
		return database.SubjectLists{}, errors.New(err.Error())
	}

	return resultSubjectID, nil
}
