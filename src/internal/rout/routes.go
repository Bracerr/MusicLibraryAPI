package rout

import (
	"MusicLibraryAPI/src/internal/delivery/hhtp"
	"MusicLibraryAPI/src/internal/service"
	"github.com/labstack/echo/v4"
)

func SetupSongRoutes(e *echo.Echo, songService *service.SongService) {
	songHandler := hhtp.NewSongHandler(songService)
	songGroup := e.Group("/songs")

	songGroup.GET("", songHandler.GetSongs)
	songGroup.GET("/text", songHandler.GetSongText)
	songGroup.DELETE("", songHandler.DeleteSong)
	songGroup.PUT("", songHandler.UpdateSong)
}
