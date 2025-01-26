package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/A5TA/NutriTrack-Ai-backend/internal/config"
	"github.com/A5TA/NutriTrack-Ai-backend/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	//Get the Env Variables Loaded
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Initialize The MongoDB Client
	config.InitializeMongoClient()

	httpPort := os.Getenv("API_PORT")
	if httpPort == "" {
		httpPort = "8050"
	}

	r := router.New() // Initialize router

	// Create server with timeout - https://github.com/gin-gonic/examples/blob/master/secure-web-app/main.go
	srv := &http.Server{
		Addr:    "0.0.0.0:" + httpPort,
		Handler: r,
		// set timeout due CWE-400 - Potential Slowloris Attack
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("Server is Listening on Port: ", httpPort)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Failed to Start Server: ", err)
	}
}
