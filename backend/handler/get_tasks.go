package handler

import (
	"backend/gen/models"
	"backend/gen/restapi/operations/task_api"
	"backend/handler/auth_jwt"
	"backend/handler/database"
	"backend/handler/dotenv"
	"errors"
	"strconv"

	"github.com/go-openapi/runtime/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	taskListsFromDB, err := GetTasksByUserIdFromDB(userID, isArchived)
	if err != nil {
		return task_api.NewGetTasksServiceUnavailable()
	}

	// 結果を入れる変数
	result := models.TasksMultiple{}

	// 結果を詰めていく
	length := len(taskListsFromDB)
	result.TotalCount = int64(length)
	for i := 0; i < length; i++ {
		// tmp変数
		tmpTaskSingle := models.TaskSingle{}
		tmpTaskSingle.CreatedAt = strconv.FormatInt(taskListsFromDB[i].CreatedAt.Unix(), 10)
		tmpTaskSingle.DeadlineAt = strconv.FormatInt(taskListsFromDB[i].DeadlineAt.Unix(), 10)
		tmpTaskSingle.IsArchived = taskListsFromDB[i].IsArchived
		tmpTaskSingle.SubjectID = int64(taskListsFromDB[i].SubjectID)
		// TODO: subjectNameを取得・DB側でJOINの必要あり
		// tmpTaskSingle.SubjectName = taskListsFromDB[i]
		tmpTaskSingle.TaskDescription = taskListsFromDB[i].TaskDescription
		tmpTaskSingle.TaskID = int64(taskListsFromDB[i].TaskID)
		tmpTaskSingle.TaskName = taskListsFromDB[i].TaskName
		// tmp変数に入れた結果をresultに格納
		result.Tasks = append(result.Tasks, &tmpTaskSingle)
	}

	return task_api.NewGetTasksOK().WithPayload(&result)
}

func GetTasksByUserIdFromDB(userID int64, isArchived bool) ([]database.TaskLists, error){
		// 結果を入れるtaskLists
		taskLists := []database.TaskLists{}

		// dsnをdotenvから読み込み
		dsnLoadByDotenv, err := dotenv.LoadDataSourceName()
		if err != nil {
			return taskLists, errors.New("failed to read secret key")
		}
		dsn := dsnLoadByDotenv.(string)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return taskLists, errors.New(err.Error())
		}

		// SELECT文
		db.Model(&database.TaskLists{}).Where("user_id = ? AND is_archived = ?", userID, isArchived).Find(&taskLists)

		return taskLists, nil
}
