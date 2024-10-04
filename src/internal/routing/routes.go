package routing

import (
	"MusicLibraryAPI/src/internal/delivery/http"
	"MusicLibraryAPI/src/internal/service"
	"github.com/labstack/echo/v4"
)

func SetupSongRoutes(e *echo.Echo, songService *service.SongService) {
	songHandler := http.NewSongHandler(songService)

	e.GET("", songHandler.GetSongs)
	e.GET("/text", songHandler.GetSongText)
	e.DELETE("", songHandler.DeleteSong)
	e.PUT("", songHandler.UpdateSong)
	e.POST("", songHandler.CreateSong)
	e.GET("/info", songHandler.GetSongBySongAndGroup)

}
