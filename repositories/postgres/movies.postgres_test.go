package postgres

import (
	"assignment-movie/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func TestMovieGetAll(t *testing.T) {
	db, err := InitDatabase()
	if err != nil {
		panic(err)
	}
	movieRepository := NewMoviesPostgres(db)

	limit := 10
	offset := 0
	orderBy := "id ASC"
	search := "a"

	movies, total, err := movieRepository.GetAll(limit, offset, orderBy, search)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Success GetData, total data: %d\n", total)

	for i, movie := range movies {
		fmt.Printf("Movie %d: %v\n", i+1, movie)
	}
}

func TestMovieGetByID(t *testing.T) {
	db, err := InitDatabase()
	if err != nil {
		panic(err)
	}
	movieRepository := NewMoviesPostgres(db)

	id := 91
	movie, err := movieRepository.GetByID(id)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Movie : %v\n", movie)
}

func TestMovieGetByTitle(t *testing.T) {
	db, err := InitDatabase()
	if err != nil {
		panic(err)
	}

	movieRepository := NewMoviesPostgres(db)

	title := "Title"

	movie, err := movieRepository.GetByTitle(title)
	if err != nil {
		fmt.Printf("Movie Not Found : %v\n", movie)
	}

	fmt.Printf("Movie : %v\n", movie)
}

func TestMovieCheckIfExists(t *testing.T) {
	db, err := InitDatabase()
	if err != nil {
		panic(err)
	}
	movieRepository := NewMoviesPostgres(db)
	movie := models.Movies{
		ID:    91,
		Title: "Title",
	}

	tof, err := movieRepository.CheckIfExists(movie)
	if err != nil {
		panic(err)
	}

	fmt.Sprintf("Result = %t", tof)
}

func TestMovieCreate(t *testing.T) {
	db, err := InitDatabase()
	if err != nil {
		panic(err)
	}
	movieRepository := NewMoviesPostgres(db)
	movie := models.Movies{
		ID:          91,
		Title:       "Title",
		Description: "Description",
		Rating:      9.5,
		Image:       "image1000.png",
		CreatedAt:   time.Now(),
	}

	err = movieRepository.Create(movie)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Create Movie!")
}

func TestMovieUpdate(t *testing.T) {
	db, err := InitDatabase()
	if err != nil {
		panic(err)
	}

	movieRepository := NewMoviesPostgres(db)
	timeNow := time.Now()

	movie := models.Movies{
		ID:          91,
		Title:       "Title Update",
		Description: "Description Update",
		Rating:      9.5,
		Image:       "updateimage1000.png",
		UpdatedAt:   &timeNow,
	}

	err = movieRepository.Update(movie)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Update Movie!")
}

func TestMovieDelete(t *testing.T) {
	db, err := InitDatabase()
	if err != nil {
		panic(err)
	}

	movieRepository := NewMoviesPostgres(db)
	timeNow := time.Now()

	movie := models.Movies{
		ID:        91,
		DeletedAt: &timeNow,
	}

	err = movieRepository.Delete(movie)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Delete Movie!")
}
