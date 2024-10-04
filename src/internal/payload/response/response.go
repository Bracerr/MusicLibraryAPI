package response

import "time"

type SongResponse struct {
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type SongTextResponse struct {
	Song        string    `json:"song"`
	Group       string    `json:"group"`
	ReleaseDate time.Time `json:"release_date"`
	Verses      []string  `json:"verses"`
	Page        int       `json:"page"`
	Limit       int       `json:"limit"`
}

type SongsResponse struct {
	Total int            `json:"total"`
	Page  int            `json:"page"`
	Limit int            `json:"limit"`
	Songs []SongResponse `json:"songs"`
}
