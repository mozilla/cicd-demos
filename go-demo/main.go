package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
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
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the cicd demo go web app! Visit /__version__ for version information. Adding test string here"))
	})

	port := "8000"
	println("Server running on port:", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Println("Error starting server:", err)
	}
}
