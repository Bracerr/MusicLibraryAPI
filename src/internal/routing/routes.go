package routing

import (
	"MusicLibraryAPI/src/internal/delivery/http"
	"MusicLibraryAPI/src/internal/service"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func SetupSongRoutes(e *echo.Echo, songService *service.SongService) {
	songHandler := http.NewSongHandler(songService)

	e.GET("", songHandler.GetSongs)
	e.GET("/text", songHandler.GetSongText)
	e.DELETE("", songHandler.DeleteSong)
	e.PUT("", songHandler.UpdateSong)
	e.POST("", songHandler.CreateSong)
	e.GET("/info", songHandler.GetSongBySongAndGroup)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

}
