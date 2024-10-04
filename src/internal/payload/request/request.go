package request

type UpdateSongRequest struct {
	GroupName   *string `json:"group_name"`
	Song        *string `json:"song"`
	ReleaseDate *string `json:"release_date"`
	Text        *string `json:"text"`
	Link        *string `json:"link"`
}

type AddSongRequest struct {
	Group string `json:"group" validate:"required"`
	Song  string `json:"song" validate:"required"`
}
