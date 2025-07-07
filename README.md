# 📺 Video Advertisement Tracker (Golang + Kafka + PostgreSQL)

A scalable and resilient backend service to manage, track, and analyze video advertisement clicks in real time.

---

## 🚀 Features

- 🔍 List video ads via REST API.
- 🖱 Track ad clicks asynchronously using Kafka.
- 📈 Real-time analytics (click counts, CTR).
- 🧱 PostgreSQL for persistent storage.
- 🔁 Kafka-based decoupled architecture for resilience.
- 🐳 Fully containerized via Docker & Docker Compose.
- 📊 Prometheus-ready metrics endpoint (optional).

---

## 📦 Tech Stack

- **Golang** — Core application logic
- **PostgreSQL** — Ad & Click data storage
- **Kafka** — Asynchronous event processing
- **Docker & Compose** — Development & deployment
- **Prometheus (optional)** — Monitoring & metrics

---

## 🛠️ Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/ads-tracker.git
cd ads-tracker

Build & Run (Dockerized)

docker-compose up -d --build

📖 API Documentation
🔹 GET /ads

Fetches a list of available ads.
✅ Response

[
  {
    "id": 1,
    "image_url": "https://example.com/image1.jpg",
    "target_url": "https://example.com"
  },
  ...
]

🔹 POST /ads/click

Records a click event. Processes the request asynchronously using Kafka.
📤 Request

{
  "ad_id": 1,
  "timestamp": "2025-07-07T06:45:00Z",
  "ip": "192.168.1.1",
  "playback_time": 15
}

✅ Response

{
  "message": "Click event received"
}

🔹 GET /ads/analytics

Fetches real-time or near real-time ad performance metrics.
🔸 Optional Query Params

    ?minutes=10 → clicks in the last 10 minutes (default: all-time)

✅ Response

[
  {
    "ad_id": 1,
    "clicks": 120,
    "ctr": 0.045
  },
  ...
]

🧠 Concurrency & Processing Flow

The system uses Kafka for asynchronously processing click events.
🔁 Flow Summary

    API receives a click → publishes to Kafka.

    Consumer service processes Kafka messages and inserts into DB.

    Analytics API queries aggregated data directly from DB.

This ensures:

    Low latency for users

    No data loss (even under DB or API failures)

    High scalability under spikes