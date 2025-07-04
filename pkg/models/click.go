package models

import "time"

type ClickEvent struct {
	AdID             string    `json:"ad_id"`
	Timestamp        time.Time `json:"timestamp"`
	IP               string    `json:"ip"`
	VideoPlaybackSec int       `json:"video_playback_time"`
}
