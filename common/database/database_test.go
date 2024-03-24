package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func TestOpenConnection(t *testing.T) {
	dsn := "host=localhost user=postgres password=1234 dbname=movies port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	if DB == nil {
		t.Fatal("DB connection is nil")
	}
}
