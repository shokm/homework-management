package handler

import (
	"backend/gen/models"
	"backend/gen/restapi/operations/task_api"
	"backend/handler/auth_jwt"
	"backend/handler/database"
	"errors"
	"strconv"
	"time"

	"github.com/go-openapi/runtime/middleware"
)

func PostTaskByTaskID(params task_api.PostTaskByTaskIDParams, principal interface{}) middleware.Responder {

	// 認証情報からユーザーIDを取得
	header := params.HTTPRequest.Header.Get("Authorization")
	returnValue, err := auth_jwt.ValidateTokenHandler(header)
	if err != nil {
		return task_api.NewGetTaskByTaskIDUnauthorized()
	}
	userInfo := returnValue.(models.AuthReturnUser)
	userID := userInfo.UserID

	// DBからの結果を入れる構造体
	resultListFromDB := database.TaskApiResultList{}

	// 結果を入れる構造体
	result := models.TaskSingle{}

	var errorValue (error)

	// 更新と追加のAPIを共有しているため、条件分岐を行う
	resultTaskIDCount, err := database.GetCountTaskByTaskID(userID, int64(params.TaskID))
	if err != nil {
		return task_api.NewPostTaskByTaskIDInternalServerError()
	}
	if resultTaskIDCount == 0 {
		// IDが存在しない場合
		resultListFromDB, errorValue = addTaskValue(*params.Body, userID)
	} else {
		// IDが存在する場合
		resultListFromDB, errorValue = updTaskValue(*params.Body, userID, int64(params.TaskID))
	}

	if errorValue != nil {
		return task_api.NewPostTaskByTaskIDInternalServerError()
	}

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

	return task_api.NewPostTaskByTaskIDOK().WithPayload(&result)
}

// 値を追加する場合
func addTaskValue(postData models.TaskSingle, userID int64) (database.TaskApiResultList, error) {

	// 書き込む値を入れる変数
	postTaskList := database.TaskLists{}

	postTaskList.DeadlineAt, _ = time.Parse("2006-01-02 15:04:05", postData.DeadlineAt)
	postTaskList.IsArchived = postData.IsArchived
	postTaskList.StateID = uint(postData.StateID)
	postTaskList.SubjectID = uint(postData.SubjectID)
	postTaskList.TaskName = postData.TaskName
	postTaskList.TaskDescription = postData.TaskDescription
	postTaskList.UserID = uint(userID)

	// DBに書き込み
	resultTaskID, err := database.PostTaskByTaskID(postTaskList)
	if err != nil {
		return database.TaskApiResultList{}, errors.New(err.Error())
	}

	// DBから結果を取得
	resultListFromDB, err := database.GetTaskByTaskID(userID, resultTaskID)
	if err != nil {
		return database.TaskApiResultList{}, errors.New(err.Error())
	}

	return resultListFromDB, nil
}

// 値を更新する場合
func updTaskValue(postData models.TaskSingle, userID int64, taskID int64) (database.TaskApiResultList, error) {

	// 書き込む値を入れる変数
	postTaskList := database.TaskLists{}

	postTaskList.DeadlineAt, _ = time.Parse("2006-01-02 15:04:05", postData.DeadlineAt)
	postTaskList.IsArchived = postData.IsArchived
	postTaskList.StateID = uint(postData.StateID)
	postTaskList.SubjectID = uint(postData.SubjectID)
	postTaskList.TaskName = postData.TaskName
	postTaskList.TaskDescription = postData.TaskDescription

	// DBを更新
	resultTaskID, err := database.PutTaskByTaskID(postTaskList, userID, taskID)
	if err != nil {
		return database.TaskApiResultList{}, errors.New(err.Error())
	}

	return resultTaskID, nil
}
