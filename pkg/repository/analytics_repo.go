package repository

import (
	"ads-tracker/pkg/db"
	"ads-tracker/pkg/models"
	"fmt"
)

func GetAdAnalytics(minutes int) ([]models.AdAnalytics, error) {
	query := `
		SELECT ad_id, COUNT(*) as clicks
		FROM ad_clicks
		WHERE timestamp >= NOW() - INTERVAL '%d minutes'
		GROUP BY ad_id
	`
	sql := fmt.Sprintf(query, minutes)

	rows, err := db.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.AdAnalytics

	for rows.Next() {
		var a models.AdAnalytics
		if err := rows.Scan(&a.AdID, &a.Clicks); err != nil {
			return nil, err
		}
		// Fake CTR (e.g., 0.001 * clicks)
		a.CTR = float64(a.Clicks) * 0.001
		result = append(result, a)
	}

	return result, nil
}
