package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	pgDB        *sql.DB
)

func main() {
	// ---- init Postgres ----
	pgDSN := os.Getenv("POSTGRES_DSN")
	if pgDSN != "" {
		db, err := sql.Open("postgres", pgDSN)
		if err != nil {
			log.Println("Postgres init error:", err)
		} else {
			// Optional: tighten default connection settings
			db.SetMaxOpenConns(5)
			db.SetMaxIdleConns(5)
			db.SetConnMaxLifetime(5 * time.Minute)
			pgDB = db
		}
	} else {
		log.Println("POSTGRES_DSN not set; postgres health checks will return down")
	}

	// ---- init Redis ----
	redisAddr := os.Getenv("REDIS_ADDR") // e.g. "redis:6379" in Compose
	if redisAddr == "" {
		log.Println("REDIS_ADDR not set; redis health checks will return down")
	} else {
		redisClient = redis.NewClient(&redis.Options{Addr: redisAddr})
	}

	http.HandleFunc("/__version__", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Read version.json file
		versionFile, err := os.Open("/app/version.json")
		if err != nil {
			http.Error(w, "Failed to read version.json", http.StatusInternalServerError)
			log.Println("Error reading version.json:", err)
			return
		}
		defer versionFile.Close()

		// Read the content of the file
		data, err := io.ReadAll(versionFile)
		if err != nil {
			http.Error(w, "Failed to read version.json content", http.StatusInternalServerError)
			log.Println("Error reading version.json content:", err)
			return
		}

		// Write the content to the response
		w.Write(data)
	})

	http.HandleFunc("/__lbheartbeat__", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/__heartbeat__", func(w http.ResponseWriter, r *http.Request) {
		type resp struct {
			Status   string `json:"status"`   // "ok" if both up, else "degraded"
			Redis    string `json:"redis"`    // "up"/"down"
			Postgres string `json:"postgres"` // "up"/"down"
		}

		// Short timeout so this endpoint is fast and resilient
		ctx, cancel := context.WithTimeout(r.Context(), 500*time.Millisecond)
		defer cancel()

		redisStatus := "down"
		pgStatus := "down"

		// Redis ping
		if redisClient != nil {
			err := redisClient.Ping(ctx).Err()
			if err != nil {
				log.Print(err)
			} else {
				redisStatus = "up"
			}
		}

		// Postgres ping
		if pgDB != nil {
			err := pgDB.PingContext(ctx)
			if err != nil {
				log.Print(err)
			} else {
				pgStatus = "up"
			}
		}

		out := resp{
			Status:   "ok",
			Redis:    redisStatus,
			Postgres: pgStatus,
		}
		if redisStatus != "up" || pgStatus != "up" {
			out.Status = "degraded"
			w.WriteHeader(http.StatusServiceUnavailable)
		} else {
			w.WriteHeader(http.StatusOK)
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(out)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the cicd demo go web app! Visit /__version__ for version information. Use the preview label to provision an ephemeral environment for testing!, Work Week Demo!"))
	})

	port := "8000"
	log.Println("Server running on port:", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Println("Error starting server:", err)
	}
}
