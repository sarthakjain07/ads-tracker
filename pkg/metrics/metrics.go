package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ClicksProcessed = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "ad_clicks_processed_total",
			Help: "Total clicks processed",
		})

	ClickFailures = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "ad_clicks_failed_total",
			Help: "Total clicks failed to be processed",
		})
)

func Init() {
	prometheus.MustRegister(ClicksProcessed)
	prometheus.MustRegister(ClickFailures)
}
