package repository

import (
	"MusicLibraryAPI/src/internal/domain"
	"gorm.io/gorm"
)

type SongRepository struct {
	db *gorm.DB
}

func NewSongRepository(db *gorm.DB) *SongRepository {
	return &SongRepository{db: db}
}

func (r *SongRepository) GetSongs(group, song string, startDate, endDate string, page, limit int) ([]domain.Song, int64, error) {
	var songs []domain.Song
	var total int64

	query := r.db.Model(&domain.Song{})

	if group != "" {
		query = query.Where("group_name = ?", group)
	}
	if song != "" {
		query = query.Where("song LIKE ?", "%"+song+"%")
	}
	if startDate != "" {
		query = query.Where("release_date >= ?", startDate)
	}
	if endDate != "" {
		query = query.Where("release_date <= ?", endDate)
	}

	query.Count(&total)
	query.Offset((page - 1) * limit).Limit(limit).Find(&songs)

	return songs, total, query.Error
}

func (r *SongRepository) GetSongByName(song string) (*domain.Song, error) {
	var songResult domain.Song

	err := r.db.Model(&domain.Song{}).
		Where("song = ?", song).
		First(&songResult).Error

	if err != nil {
		return nil, err
	}

	return &songResult, nil
}

func (r *SongRepository) DeleteSongByName(song string) error {
	result := r.db.Where("song = ?", song).Delete(&domain.Song{})
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

func (r *SongRepository) UpdateSong(song domain.Song) error {
	return r.db.Save(&song).Error
}

func (r *SongRepository) InsertSong(song domain.Song) error {
	return r.db.Create(&song).Error
}

func (r *SongRepository) GetSongBySongAndGroup(song, group string) (*domain.Song, error) {
	var songResult domain.Song

	err := r.db.Model(&domain.Song{}).
		Where("group_name = ? AND song = ?", group, song).
		First(&songResult).Error

	if err != nil {
		return nil, err
	}

	return &songResult, nil
}
