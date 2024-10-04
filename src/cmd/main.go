package main

import (
	_ "MusicLibraryAPI/docs" // Импортируйте сгенерированные документы
	"MusicLibraryAPI/src/internal/repository"
	"MusicLibraryAPI/src/internal/routing"
	"MusicLibraryAPI/src/internal/service"
	"MusicLibraryAPI/src/pkg/configs"
	"MusicLibraryAPI/src/pkg/database"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log"
)

func main() {
	dbModel := configs.GetDbParams()
	db := database.NewClient(dbModel)
	log.Println("Connecting to database successful")

	database.ApplyMigrations(dbModel)

	songRepo := repository.NewSongRepository(db)
	songService := service.NewSongService(songRepo)

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	routing.SetupSongRoutes(e, songService)
	e.Logger.Fatal(e.Start(configs.GetServerParams().ServerHost))
}
