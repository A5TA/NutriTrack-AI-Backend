package main

import (
	"log"
	"net/http"
	"os"

	"github.com/A5TA/NutriTrack-Ai-backend/internal/ai"
	"github.com/A5TA/NutriTrack-Ai-backend/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	//Get the Env Variables Loaded
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the ONNX model
	if err := ai.InitModel(os.Getenv("Model_Path")); err != nil {
		log.Fatalf("Error initializing model: %v", err)
	}

	r := router.New()

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Failed to Start Server: ", err)
	}
}
