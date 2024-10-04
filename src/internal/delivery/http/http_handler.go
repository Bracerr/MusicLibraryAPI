package http

import (
	"MusicLibraryAPI/src/internal/payload/request"
	"MusicLibraryAPI/src/internal/payload/response"
	"MusicLibraryAPI/src/internal/service"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

type SongHandler struct {
	service *service.SongService
}

func validateStruct(s interface{}) error {
	validate := validator.New()
	return validate.Struct(s)
}

func NewSongHandler(service *service.SongService) *SongHandler {
	return &SongHandler{service: service}
}

func (h *SongHandler) GetSongs(c echo.Context) error {
	group := c.QueryParam("group")
	song := c.QueryParam("song")
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")
	pageStr := c.QueryParam("page")
	limitStr := c.QueryParam("limit")

	page := 1
	limit := 10

	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil {
			page = p
		}
	}

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = l
		}
	}

	songs, total, err := h.service.GetSongs(group, song, startDate, endDate, page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	var songResponses []response.SongResponse
	for _, s := range songs {
		songResponses = append(songResponses, response.SongResponse{
			Group:       s.GroupName,
			Song:        s.Song,
			ReleaseDate: s.ReleaseDate.Format("2006-01-02"),
			Text:        s.Text,
			Link:        s.Link,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"total": total,
		"page":  page,
		"limit": limit,
		"songs": songResponses,
	})
}

// GetSongText godoc
// @Title GetUser
// @Description Get user information
// @Accept  json
// @Param nick_name formData string true "nickname"
// @Param user_name formData string true "user name"
// @Param password formData string true "password"
// @Param age formData int true "age"
// @Success 200 "Successfully obtained information"
// @Failure 400 "Failed to obtain information"
// @Router /getUser [get]
func (h *SongHandler) GetSongText(c echo.Context) error {
	songName := c.QueryParam("song")
	pageStr := c.QueryParam("page")
	limitStr := c.QueryParam("limit")

	page := 1
	limit := 1

	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil {
			page = p
		}
	}

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = l
		}
	}

	song, err := h.service.GetSongByName(songName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if song.Song != songName {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Song not found"})
	}
	verses := strings.Split(song.Text, "\\n")

	start := (page - 1) * limit
	end := start + limit

	if start >= len(verses) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "No verses found"})
	}

	if end > len(verses) {
		end = len(verses)
	}

	paginatedVerses := verses[start:end]

	return c.JSON(http.StatusOK, map[string]interface{}{
		"song":         song.Song,
		"group":        song.GroupName,
		"release_date": song.ReleaseDate,
		"verses":       paginatedVerses,
		"page":         page,
		"limit":        limit,
	})
}

func (h *SongHandler) DeleteSong(c echo.Context) error {
	songName := c.QueryParam("song")

	err := h.service.DeleteSongByName(songName)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Song not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Song deleted successfully"})
}

func (h *SongHandler) UpdateSong(c echo.Context) error {
	var updateRequest request.UpdateSongRequest
	if err := c.Bind(&updateRequest); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	songName := c.QueryParam("song")
	err := h.service.UpdateSongByName(songName, updateRequest)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "Song not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Song updated successfully"})
}

func (h *SongHandler) CreateSong(c echo.Context) error {
	var addSongRequest request.AddSongRequest
	if err := c.Bind(&addSongRequest); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	err := validateStruct(addSongRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	err = h.service.CreateSong(addSongRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Song created successfully"})
}

func (h *SongHandler) GetSongBySongAndGroup(c echo.Context) error {
	songName := c.QueryParam("song")
	groupName := c.QueryParam("group")
	song, err := h.service.GetSongBySongAndGroup(songName, groupName)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, song)

}
