package request

type UpdateSongRequest struct {
	GroupName   *string `json:"group_name"`
	Song        *string `json:"song"`
	ReleaseDate *string `json:"release_date"` // Измените на *string
	Text        *string `json:"text"`
	Link        *string `json:"link"`
}
