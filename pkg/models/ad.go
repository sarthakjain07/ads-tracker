package models

type Ad struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	ImageURL  string `json:"image_url"`
	TargetURL string `json:"target_url"`
}
