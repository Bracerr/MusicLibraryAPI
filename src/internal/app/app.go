package app

import (
	"MusicLibraryAPI/src/internal/repository"
	"MusicLibraryAPI/src/internal/rout"
	"MusicLibraryAPI/src/internal/service"
	"MusicLibraryAPI/src/pkg/configs"
	"MusicLibraryAPI/src/pkg/database"
	"github.com/labstack/echo/v4"
)

func Run() {
	dbModel := configs.GetDbParams()
	db := database.NewClient(dbModel)

	database.ApplyMigrations(dbModel)

	songRepo := repository.NewSongRepository(db)
	songService := service.NewSongService(songRepo)

	e := echo.New()

	rout.SetupSongRoutes(e, songService)
	e.Logger.Fatal(e.Start(configs.GetServerParams().ServerHost))
}
