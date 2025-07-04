package repository

import (
	"ads-tracker/pkg/db"
	"ads-tracker/pkg/models"
	"fmt"
)

func GetAllAds() ([]models.Ad, error) {
	rows, err := db.DB.Query("SELECT id, title, image_url, target_url FROM ads")
	if err != nil {
		fmt.Println("DB Query Error:", err)
		return nil, err
	}

	defer rows.Close()

	var ads []models.Ad
	for rows.Next() {
		var ad models.Ad
		if err := rows.Scan(&ad.ID, &ad.Title, &ad.ImageURL, &ad.TargetURL); err != nil {
			return nil, err
		}
		ads = append(ads, ad)
	}

	return ads, nil
}
