package database

import (
	"backend/gen/models"
	"backend/handler/dotenv"
	"errors"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TaskApiResultList struct {
	TaskID          uint
	UserID          uint
	SubjectID       uint
	SubjectName		string
	StateID         uint
	TaskName        string
	TaskDescription string
	IsArchived      bool
	DeadlineAt      time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func GetTasksByUserID(userID int64, isArchived bool) ([]TaskApiResultList, error){
	// 結果を入れるtaskLists
	resultLists := []TaskApiResultList{}

	// dsnをdotenvから読み込み
	dsnLoadByDotenv, err := dotenv.LoadDataSourceName()
	if err != nil {
		return resultLists, errors.New("failed to read secret key")
	}
	dsn := dsnLoadByDotenv.(string)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return resultLists, errors.New(err.Error())
	}

	// SELECT文
	// db.Model(&TaskLists{}).Where("task_lists.user_id = ? AND task_lists.is_archived = ?", userID, isArchived).Joins("left join subject_lists on task_lists.subject_id = subject_lists.subject_id").Find(&resultLists)
	// TODO: db.Debug()
	db.Debug().Model(&TaskLists{}).Select("subject_lists.subject_name, task_lists.task_id, task_lists.user_id, task_lists.subject_id, task_lists.state_id, task_lists.task_name, task_lists.task_description, task_lists.is_archived, task_lists.deadline_at, task_lists.created_at, task_lists.updated_at").Where("task_lists.user_id = ? AND task_lists.is_archived = ?", userID, isArchived).Joins("left join subject_lists on task_lists.subject_id = subject_lists.subject_id").Order("task_lists.task_id asc").Scan(&resultLists)

	// セッションを切る
	sqlDB, err := db.DB()
	if err != nil {
		return resultLists, errors.New(err.Error())
	}
	sqlDB.Close()

	return resultLists, nil
}

func GetTaskByTaskID(userID int64, taskID int64) (TaskApiResultList, error){
	// 結果を入れるtaskLists
	resultList := TaskApiResultList{}

	// dsnをdotenvから読み込み
	dsnLoadByDotenv, err := dotenv.LoadDataSourceName()
	if err != nil {
		return resultList, errors.New("failed to read secret key")
	}
	dsn := dsnLoadByDotenv.(string)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return resultList, errors.New(err.Error())
	}

	// SELECT文
	// db.Model(&TaskLists{}).Where("task_lists.user_id = ? AND task_lists.task_id = ?", userID, taskID).Joins("left join subject_lists on task_lists.subject_id = subject_lists.subject_id").First(&resultList)
	// TODO: db.Debug()
	db.Debug().Model(&TaskLists{}).Where("task_lists.user_id = ? AND task_lists.task_id = ?", userID, taskID).Joins("left join subject_lists on task_lists.subject_id = subject_lists.subject_id").First(&resultList)

	// セッションを切る
	sqlDB, err := db.DB()
	if err != nil {
		return resultList, errors.New(err.Error())
	}
	sqlDB.Close()

	return resultList, nil
}

func GetCountTaskByTaskID(userID int64, taskID int64) (int64, error){
	// 結果を入れるresultCount
	var count (int64)

	// dsnをdotenvから読み込み
	dsnLoadByDotenv, err := dotenv.LoadDataSourceName()
	if err != nil {
		return count, errors.New("failed to read secret key")
	}
	dsn := dsnLoadByDotenv.(string)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return count, errors.New(err.Error())
	}

	// db.Model(&TaskLists{}).Where("task_id = ? AND user_id", taskID, userID).Count(&count)
	// TODO: db.Debug()
	db.Debug().Model(&TaskLists{}).Where("task_id = ? AND user_id = ?", taskID, userID).Count(&count)

	// セッションを切る
	sqlDB, err := db.DB()
	if err != nil {
		return count, errors.New(err.Error())
	}
	sqlDB.Close()

	return count, nil
}

func PostTaskByTaskID(postTaskLists TaskLists) (int64, error){
	// dsnをdotenvから読み込み
	dsnLoadByDotenv, err := dotenv.LoadDataSourceName()
	if err != nil {
		return 0, errors.New("failed to read secret key")
	}
	dsn := dsnLoadByDotenv.(string)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return 0, errors.New(err.Error())
	}

	// db.Create(postTaskLists)
	// TODO: db.Debug()
	db.Debug().Create(&postTaskLists)

	// セッションを切る
	sqlDB, err := db.DB()
	if err != nil {
		return 0, errors.New(err.Error())
	}
	sqlDB.Close()

	return int64(postTaskLists.TaskID), nil
}

func PutTaskByTaskID(postTaskLists TaskLists, userID int64, taskID int64) (TaskApiResultList, error){
	// 結果を入れるtaskLists
	resultList := TaskApiResultList{}

	// dsnをdotenvから読み込み
	dsnLoadByDotenv, err := dotenv.LoadDataSourceName()
	if err != nil {
		return resultList, errors.New("failed to read secret key")
	}
	dsn := dsnLoadByDotenv.(string)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return resultList, errors.New(err.Error())
	}

	// Update
	// db.Model(&TaskLists{}).Where("task_id = ? AND user_id = ?", taskID, userID).Updates(postTaskLists)
	// TODO: db.Debug()
	db.Debug().Model(&TaskLists{}).Where("task_id = ? AND user_id = ?", taskID, userID).Updates(postTaskLists)

	// Setect
	db.Model(&TaskLists{}).Where("task_lists.user_id = ? AND task_lists.task_id = ?", userID, taskID).Joins("left join subject_lists on task_lists.subject_id = subject_lists.subject_id").First(&resultList)

	// セッションを切る
	sqlDB, err := db.DB()
	if err != nil {
		return resultList, errors.New(err.Error())
	}
	sqlDB.Close()

	return resultList, nil
}

func GetSubjectsByUserID(userID int64, isArchived bool) ([]models.SubjectSingle, error){
	// 結果を入れるtaskLists
	resultLists := []models.SubjectSingle{}

	// dsnをdotenvから読み込み
	dsnLoadByDotenv, err := dotenv.LoadDataSourceName()
	if err != nil {
		return resultLists, errors.New("failed to read secret key")
	}
	dsn := dsnLoadByDotenv.(string)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return resultLists, errors.New(err.Error())
	}

	// SELECT文
	db.Model(&SubjectLists{}).Where("user_id = ? AND is_archived = ?", userID, isArchived).Order("subject_id asc").Find(&resultLists)

	// セッションを切る
	sqlDB, err := db.DB()
	if err != nil {
		return resultLists, errors.New(err.Error())
	}
	sqlDB.Close()

	return resultLists, nil
}

func GetSubjectBySubjectID(userID int64, SubjectID int64) (SubjectLists, error){
	// 結果を入れるtaskLists
	resultList := SubjectLists{}

	// dsnをdotenvから読み込み
	dsnLoadByDotenv, err := dotenv.LoadDataSourceName()
	if err != nil {
		return resultList, errors.New("failed to read secret key")
	}
	dsn := dsnLoadByDotenv.(string)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return resultList, errors.New(err.Error())
	}

	// SELECT文
	db.Model(&SubjectLists{}).Where("user_id = ? AND subject_id = ?", userID, SubjectID).First(&resultList)

	// セッションを切る
	sqlDB, err := db.DB()
	if err != nil {
		return resultList, errors.New(err.Error())
	}
	sqlDB.Close()

	return resultList, nil
}

func GetCountSubjectBySubjectID(userID int64, SubjectID int64) (int64, error){
	// 結果を入れるresultCount
	var count (int64)

	// dsnをdotenvから読み込み
	dsnLoadByDotenv, err := dotenv.LoadDataSourceName()
	if err != nil {
		return count, errors.New("failed to read secret key")
	}
	dsn := dsnLoadByDotenv.(string)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return count, errors.New(err.Error())
	}

	// db.Model(&TaskLists{}).Where("task_id = ? AND user_id", SubjectID, userID).Count(&count)
	// TODO: db.Debug()
	db.Debug().Model(&SubjectLists{}).Where("subject_id = ? AND user_id = ?", SubjectID, userID).Count(&count)

	// セッションを切る
	sqlDB, err := db.DB()
	if err != nil {
		return count, errors.New(err.Error())
	}
	sqlDB.Close()

	return count, nil
}

func PostSubjectBySubjectID(postTaskLists SubjectLists) (int64, error){
	// dsnをdotenvから読み込み
	dsnLoadByDotenv, err := dotenv.LoadDataSourceName()
	if err != nil {
		return 0, errors.New("failed to read secret key")
	}
	dsn := dsnLoadByDotenv.(string)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return 0, errors.New(err.Error())
	}

	// db.Create(postTaskLists)
	// TODO: db.Debug()
	db.Debug().Create(&postTaskLists)

	// セッションを切る
	sqlDB, err := db.DB()
	if err != nil {
		return 0, errors.New(err.Error())
	}
	sqlDB.Close()

	return int64(postTaskLists.SubjectID), nil
}

func PutSubjectBySubjectID(postSubjectLists SubjectLists, userID int64, SubjectID int64) (SubjectLists, error){
	// 結果を入れるtaskLists
	resultList := SubjectLists{}

	// dsnをdotenvから読み込み
	dsnLoadByDotenv, err := dotenv.LoadDataSourceName()
	if err != nil {
		return resultList, errors.New("failed to read secret key")
	}
	dsn := dsnLoadByDotenv.(string)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return resultList, errors.New(err.Error())
	}

	// Update
	// db.Model(&TaskLists{}).Where("task_id = ? AND user_id = ?", taskID, userID).Updates(postTaskLists)
	// TODO: db.Debug()
	db.Debug().Model(&SubjectLists{}).Where("subject_id = ? AND user_id = ?", SubjectID, userID).Updates(postSubjectLists)

	// SELECT文
	db.Model(&SubjectLists{}).Where("user_id = ? AND subject_id = ?", userID, SubjectID).First(&resultList)

	// セッションを切る
	sqlDB, err := db.DB()
	if err != nil {
		return resultList, errors.New(err.Error())
	}
	sqlDB.Close()

	return resultList, nil
}

func GetTasksBySubjectID(userID int64, subjectID int64, isArchived bool) ([]TaskLists, error){
	// 結果を入れるtaskLists
	resultLists := []TaskLists{}

	// dsnをdotenvから読み込み
	dsnLoadByDotenv, err := dotenv.LoadDataSourceName()
	if err != nil {
		return resultLists, errors.New("failed to read secret key")
	}
	dsn := dsnLoadByDotenv.(string)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return resultLists, errors.New(err.Error())
	}

	// SELECT文
	// db.Model(&TaskLists{}).Where("task_lists.user_id = ? AND task_lists.is_archived = ?", userID, isArchived).Joins("left join subject_lists on task_lists.subject_id = subject_lists.subject_id").Find(&resultLists)
	// TODO: db.Debug()
	db.Debug().Model(&TaskLists{}).Where("user_id = ? AND subject_id = ? AND is_archived = ?", userID, subjectID, isArchived).Order("task_id asc").Scan(&resultLists)

	// セッションを切る
	sqlDB, err := db.DB()
	if err != nil {
		return resultLists, errors.New(err.Error())
	}
	sqlDB.Close()

	return resultLists, nil
}
