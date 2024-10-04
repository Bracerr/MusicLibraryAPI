package main

import (
	"MusicLibraryAPI/src/internal/repository"
	"MusicLibraryAPI/src/internal/routing"
	"MusicLibraryAPI/src/internal/service"
	"MusicLibraryAPI/src/pkg/configs"
	"MusicLibraryAPI/src/pkg/database"
	"github.com/labstack/echo/v4"
	_ "github.com/swaggo/echo-swagger/example/docs"
	"log"
)

// @title Swagger API
// @version 1.0
// @description This is a sample server
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:63342
// @BasePath /
func main() {
	dbModel := configs.GetDbParams()
	db := database.NewClient(dbModel)
	log.Println("Connecting to database successful")

	database.ApplyMigrations(dbModel)

	songRepo := repository.NewSongRepository(db)
	songService := service.NewSongService(songRepo)

	e := echo.New()

	routing.SetupSongRoutes(e, songService)
	e.Logger.Fatal(e.Start(configs.GetServerParams().ServerHost))
}
