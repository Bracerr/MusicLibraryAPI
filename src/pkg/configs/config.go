package configs

import (
	db "MusicLibraryAPI/src/pkg/database"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

type ServerParams struct {
	ServerHost string
}

func GetDbParams() db.DbInitModel {
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Ошибка при получении текущей директории: %v", err)
	}

	envFilePath := filepath.Join(filepath.Dir(filepath.Dir(rootDir)), ".env")

	envErr := godotenv.Load(envFilePath)
	if envErr != nil {
		log.Fatalf("Ошибка при загрузке .env файла: %v", envErr)
	}

	dbInitModel := db.DbInitModel{
		DbHost:     os.Getenv("DB_HOST"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		DbPort:     os.Getenv("DB_PORT"),
	}
	return dbInitModel
}

func GetServerParams() ServerParams {
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Ошибка при получении текущей директории: %v", err)
	}
	envFilePath := filepath.Join(filepath.Dir(filepath.Dir(rootDir)), ".env")
	envErr := godotenv.Load(envFilePath)
	if envErr != nil {
		log.Fatalf("Ошибка при загрузке .env файла: %v", envErr)
	}
	return ServerParams{ServerHost: os.Getenv("SERVER_PORT")}
}
