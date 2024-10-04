package database

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type DbInitModel struct {
	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string
	DbPort     string
}

func NewClient(dbModel DbInitModel) *gorm.DB {
	dsn := "host=" + dbModel.DbHost +
		" user=" + dbModel.DbUser +
		" password=" + dbModel.DbPassword +
		" dbname=" + dbModel.DbName +
		" port=" + dbModel.DbPort +
		" sslmode=disable"

	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	return db
}

func ApplyMigrations(dbModel DbInitModel) {
	dsn := "postgres://" + dbModel.DbUser + ":" + dbModel.DbPassword + "@" + dbModel.DbHost + ":" + dbModel.DbPort + "/" + dbModel.DbName + "?sslmode=disable"

	migrationPath := os.Getenv("MIGRATION_PATH")
	if migrationPath == "" {
		migrationPath = "file://../internal/db/migrations"
	}

	m, err := migrate.New(
		migrationPath,
		dsn,
	)

	if err != nil {
		log.Fatalf("failed to create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("Миграции не требуются: база данных уже на актуальной версии")
		} else {
			log.Fatalf("Ошибка при выполнении миграции: %v", err)
		}
	} else {
		log.Println("Миграция успешно выполнена")
	}

	if err, _ := m.Close(); err != nil {
		log.Fatalf("Ошибка при закрытии миграций: %v", err)
	}

}
