package models

type AdAnalytics struct {
	AdID   string  `json:"ad_id"`
	Clicks int     `json:"clicks"`
	CTR    float64 `json:"ctr"`
}
