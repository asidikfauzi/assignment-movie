package database

import (
	"assignment-movie/common/helper"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func InitDatabase() (*gorm.DB, error) {
	var err error

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		helper.GetEnv("DB_HOST"),
		helper.GetEnv("DB_USERNAME"),
		helper.GetEnv("DB_PASSWORD"),
		helper.GetEnv("DB_DATABASE"),
		helper.GetEnv("DB_PORT"),
	)

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		panic("Failed Connect To Database: " + err.Error())
	}

	sqlDB, err := DB.DB()
	if err != nil {
		panic("Failed to access database connection pool: " + err.Error())
	}

	sqlDB.SetConnMaxIdleTime(10 * time.Second)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)

	return DB, err
}
