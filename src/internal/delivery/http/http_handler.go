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

// GetSongs godoc
// @Summary Get a list of songs
// @Description Retrieve a list of songs based on optional filters such as group name, song name, release date range, and pagination.
// @Accept  json
// @Produce  json
// @Param group query string false "Group name to filter songs"
// @Param song query string false "Song name to filter songs"
// @Param start_date query string false "Start date to filter songs (format: YYYY-MM-DD)"
// @Param end_date query string false "End date to filter songs (format: YYYY-MM-DD)"
// @Param page query int false "Page number for pagination (default is 1)"
// @Param limit query int false "Number of songs per page (default is 10)"
// @Success 200 {object} response.SongsResponse "Successfully retrieved songs"
// @Failure 500 {object} map[string]string "Internal server error, returns an object with an 'error' field"
// @Router / [get]
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

	SongsResponse := response.SongsResponse{
		Total: int(total),
		Page:  page,
		Limit: limit,
		Songs: songResponses,
	}

	return c.JSON(http.StatusOK, SongsResponse)
}

// GetSongText godoc
// @Summary Get the text of a song
// @Description Retrieve the text of a song by its name, with optional pagination for verses. 1 page= 1 verse
// @Accept  json
// @Produce  json
// @Param song query string true "Name of the song to retrieve"
// @Param page query int false "Page number for pagination (default is 1)"
// @Param limit query int false "Number of verses per page (default is 1)"
// @Success 200 {object} response.SongTextResponse "Successfully retrieved song text"
// @Failure 400 {object} map[string]string "Bad request, invalid parameters"
// @Failure 404 {object} map[string]string "Song not found or no verses found"
// @Failure 500 {object} map[string]string "Internal server error, returns an object with an 'error' field"
// @Router /text [get]
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

	textResponse := response.SongTextResponse{
		Song:        song.Song,
		Group:       song.GroupName,
		ReleaseDate: song.ReleaseDate,
		Verses:      paginatedVerses,
		Page:        page,
		Limit:       limit,
	}

	return c.JSON(http.StatusOK, textResponse)
}

// DeleteSong godoc
// @Summary Delete a song
// @Description Delete a song by its name.
// @Accept  json
// @Produce  json
// @Param song query string true "Name of the song to delete"
// @Success 200 {object} map[string]string "Successfully deleted song"
// @Failure 400 {object} map[string]string "Bad request, invalid parameters"
// @Failure 404 {object} map[string]string "Song not found"
// @Failure 500 {object} map[string]string "Internal server error, returns an object with an 'error' field"
// @Router / [delete]
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

// UpdateSong godoc
// @Summary Update a song
// @Description Update the details of a song by its name.
// @Accept  json
// @Produce  json
// @Param song query string true "Name of the song to update"
// @Param updateRequest body request.UpdateSongRequest true "Details to update the song. Only the field that was passed is updated"
// @Success 200 {object} map[string]string "Successfully updated song"
// @Failure 400 {object} map[string]string "Bad request, invalid parameters"
// @Failure 404 {object} map[string]string "Song not found"
// @Failure 500 {object} map[string]string "Internal server error, returns an object with an 'error' field"
// @Router / [put]
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

// CreateSong godoc
// @Summary Create a new song
// @Description Create a new song with the provided details.
// @Accept  json
// @Produce  json
// @Param addSongRequest body request.AddSongRequest true "Details of the song to create"
// @Success 200 {object} map[string]string "Successfully created song"
// @Failure 400 {object} map[string]string "Bad request, invalid parameters"
// @Router / [post]
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

// GetSongBySongAndGroup godoc
// @Summary Get a song by its name and group
// @Description Retrieve a song based on its name and the group it belongs to.
// @Accept  json
// @Produce  json
// @Param song query string true "Name of the song to retrieve"
// @Param group query string true "Name of the group to filter the song"
// @Success 200 {object} response.SongResponse "Successfully retrieved song"
// @Failure 400 {object} map[string]string "Bad request, invalid parameters"
// @Router /info [get]
func (h *SongHandler) GetSongBySongAndGroup(c echo.Context) error {
	songName := c.QueryParam("song")
	groupName := c.QueryParam("group")
	song, err := h.service.GetSongBySongAndGroup(songName, groupName)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, song)

}
