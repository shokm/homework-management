package database

import (
	"time"
)

type UserLists struct {
	UserID uint `gorm:"primaryKey"`
	ScreenName string `gorm:"not null"`
	Password string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SubjectLists struct {
	SubjectID uint `gorm:"primaryKey"`
	UserID uint
	SubjectName string `gorm:"not null"`
	IsArchived bool `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type StateLists struct {
	StateID uint `gorm:"primaryKey"`
	UserID uint
	StateName string `gorm:"not null"`
	IsArchived bool `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TaskLists struct {
	TaskID uint `gorm:"primaryKey"`
	UserID uint
	SubjectID uint
	StateID uint
	TaskName string `gorm:"not null"`
	TaskDescription string
	IsArchived bool `gorm:"not null"`
	DeadlineAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
