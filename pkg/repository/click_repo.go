package repository

import (
	"ads-tracker/pkg/db"
	"ads-tracker/pkg/models"
)

func SaveClickEvent(click models.ClickEvent) error {
	query := `
		INSERT INTO ad_clicks (ad_id, timestamp, ip, video_playback_sec)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (ad_id, timestamp, ip) DO NOTHING;
	`

	_, err := db.DB.Exec(query, click.AdID, click.Timestamp, click.IP, click.VideoPlaybackSec)
	return err
}

