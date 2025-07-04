// func GetAds(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("GET /ads - list of ads"))
// }

// func PostClick(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("POST /ads/click - record click"))
// }

//	func GetAnalytics(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte("GET /ads/analytics - metrics"))
//	}
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"ads-tracker/pkg/kafka"
	"ads-tracker/pkg/models"
	"ads-tracker/pkg/repository"
)

var producer *kafka.ClickProducer

func InitClickProducer(p *kafka.ClickProducer) {
	producer = p
}

// PostClick
func PostClick(w http.ResponseWriter, r *http.Request) {
	var click models.ClickEvent

	if err := json.NewDecoder(r.Body).Decode(&click); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if click.AdID == "" || click.VideoPlaybackSec < 0 {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	if click.Timestamp.IsZero() {
		click.Timestamp = time.Now()
	}
	click.IP = r.RemoteAddr

	go producer.SendClickEvent(click)

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("click received"))
}

// GetAds
func GetAds(w http.ResponseWriter, r *http.Request) {
	ads, err := repository.GetAllAds()
	if err != nil {
		http.Error(w, "failed to fetch ads", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ads)
}

// GetAnalytics
func GetAnalytics(w http.ResponseWriter, r *http.Request) {
	minutesStr := r.URL.Query().Get("minutes")
	minutes := 30 // default
	if minutesStr != "" {
		if parsed, err := strconv.Atoi(minutesStr); err == nil {
			minutes = parsed
		}
	}

	analytics, err := repository.GetAdAnalytics(minutes)
	if err != nil {
		http.Error(w, "failed to fetch analytics", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(analytics)
}
