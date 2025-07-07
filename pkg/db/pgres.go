package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitPostgres() {
	url := os.Getenv("DB_URL")

	var err error
	for i := 0; i < 10; i++ {
		DB, err = sql.Open("postgres", url)
		if err == nil && DB.Ping() == nil {
			fmt.Println("Connected to PostgreSQL")
			return
		}

		log.Printf("Failed to connect to DB (attempt %d/10): %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	log.Fatalf("Could not connect to DB after 10 attempts: %v", err)
}
