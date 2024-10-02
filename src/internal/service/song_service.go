package service

import (
	"MusicLibraryAPI/src/internal/domain"
	"MusicLibraryAPI/src/internal/repository"
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
