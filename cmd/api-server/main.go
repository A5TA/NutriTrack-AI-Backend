package main

import (
	"log"
	"net/http"

	"github.com/A5TA/NutriTrack-Ai-backend/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	//Get the Env Variables Loaded
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := router.New()

	if err := http.ListenAndServe(":8050", r); err != nil {
		log.Fatal("Failed to Start Server: ", err)
	}
}
