package helper

import (
	"fmt"
	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	var appConfig map[string]string
	appConfig, err := godotenv.Read()

	if err != nil {
		fmt.Println("Error reading .env file")
	}
	return appConfig[key]
}
