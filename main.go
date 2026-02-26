package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dayaj1222/video-streamer/streamer"
)

func main() {

	s := &streamer.Streamer{VideoDir: "videos"}

	mux := http.NewServeMux()

	mux.HandleFunc("/videos", func(w http.ResponseWriter, r *http.Request) {
		videos, err := s.GetAllVideos()
		if err != nil {
			http.Error(w, "Failed to list videos", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(videos)
	})

	mux.HandleFunc("/video", func(w http.ResponseWriter, r *http.Request) {
		filename := r.URL.Query().Get("file")
		if filename == "" {
			http.Error(w, "Missing file parameter", http.StatusBadRequest)
			return
		}
		err := s.StreamVideo(w, r, filename)
		if err != nil {
			http.Error(w, "Failed to stream video", http.StatusNotFound)
		}
	})

	addr := ":8080"
	log.Printf("Starting server on %s", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
