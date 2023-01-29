package main

import (
	"backend/handler/database"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgrespw dbname=postgres port=55000 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&database.UserLists{}, &database.SubjectLists{}, &database.StateLists{}, &database.TaskLists{})
	fmt.Println("migrated")
}
