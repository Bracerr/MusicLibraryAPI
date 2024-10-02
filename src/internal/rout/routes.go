package rout

import (
	"MusicLibraryAPI/src/internal/delivery/hhtp"
	"MusicLibraryAPI/src/internal/service"
	"github.com/labstack/echo/v4"
)

func SetupSongRoutes(e *echo.Echo, songService *service.SongService) {
	songHandler := hhtp.NewSongHandler(songService)
	e.GET("/songs", songHandler.GetSongs)
	e.GET("/songs/text", songHandler.GetSongText)
}
