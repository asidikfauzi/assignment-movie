package services

import (
	pgRepo "assignment-movie/repositories/postgres"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mime/multipart"
	"net/http/httptest"
	"testing"
	"time"
)

var DB *gorm.DB

func InitDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=1234 dbname=movies port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	var err error
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

func createDummyFileHeader() *multipart.FileHeader {
	buf := bytes.NewBuffer([]byte("dummy file content"))

	fileHeader := &multipart.FileHeader{
		Filename: "dummy.txt",
		Size:     int64(buf.Len()),
	}

	return fileHeader
}

func TestMovieGetAll(t *testing.T) {
	timeNow := time.Now()

	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
	db, err := InitDatabase()
	if err != nil {
		panic(err)
	}

	moviePostgres := pgRepo.NewMoviesPostgres(db)
	movieService := NewMoviesService(moviePostgres)

	page := "1"
	limit := "10"
	orderBy := "id ASC"
	search := "a"

	movies, total, err := movieService.GetAll(ginContext, page, limit, orderBy, search, timeNow)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Success GetData, total data: %d\n", total)

	for i, movie := range movies {
		fmt.Printf("Movie %d: %v\n", i+1, movie)
	}
}

func TestMovieGetByID(t *testing.T) {
	timeNow := time.Now()

	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
	db, err := InitDatabase()
	if err != nil {
		panic(err)
	}

	moviePostgres := pgRepo.NewMoviesPostgres(db)
	movieService := NewMoviesService(moviePostgres)

	id := 1
	movie, err := movieService.GetByID(ginContext, id, timeNow)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Movie: %v\n", movie)
}
