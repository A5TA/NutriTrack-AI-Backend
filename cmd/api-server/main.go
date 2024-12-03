package main

import (
	"log"
	"net/http"

	"github.com/A5TA/NutriTrack-Ai-backend/internal/router"
)

func main() {
	r := router.New()

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Failed to Start Server: ", err)
	}
}
