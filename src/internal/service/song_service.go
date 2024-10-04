package service

import (
	"MusicLibraryAPI/src/internal/domain"
	"MusicLibraryAPI/src/internal/payload/request"
	"MusicLibraryAPI/src/internal/repository"
	"errors"
	"time"
)

type SongService struct {
	repo *repository.SongRepository
}

func NewSongService(repo *repository.SongRepository) *SongService {
	return &SongService{repo: repo}
}

func (s *SongService) GetSongs(group, song, startDate, endDate string, page, limit int) ([]domain.Song, int64, error) {
	return s.repo.GetSongs(group, song, startDate, endDate, page, limit)
}

func (s *SongService) GetSongByName(song string) (*domain.Song, error) {
	return s.repo.GetSongByName(song)
}

func (s *SongService) DeleteSongByName(song string) error {
	return s.repo.DeleteSongByName(song)
}

func (s *SongService) UpdateSongByName(songName string, updateRequest request.UpdateSongRequest) error {
	song, err := s.repo.GetSongByName(songName)
	if err != nil {
		return err
	}

	if updateRequest.GroupName != nil {
		song.GroupName = *updateRequest.GroupName
	}
	if updateRequest.Song != nil {
		song.Song = *updateRequest.Song
	}
	if updateRequest.ReleaseDate != nil {
		releaseDate, err := time.Parse("2006-01-02", *updateRequest.ReleaseDate)
		if err != nil {
			return errors.New("invalid release date format, expected YYYY-MM-DD")
		}
		song.ReleaseDate = releaseDate
	}
	if updateRequest.Text != nil {
		song.Text = *updateRequest.Text
	}
	if updateRequest.Link != nil {
		song.Link = *updateRequest.Link
	}

	return s.repo.UpdateSong(*song)
}

func (s *SongService) CreateSong(songRequest request.AddSongRequest) error {
	song := domain.Song{
		GroupName:   songRequest.Group,
		Song:        songRequest.Song,
		ReleaseDate: time.Now(),
		Text:        "",
		Link:        "",
	}
	return s.repo.InsertSong(song)
}

func (s *SongService) GetSongBySongAndGroup(song string, group string) (*domain.Song, error) {
	return s.repo.GetSongBySongAndGroup(song, group)
}
