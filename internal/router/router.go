package router

import (
	"github.com/A5TA/NutriTrack-Ai-backend/internal/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// New creates a new Gin router and sets up routes
func New() *gin.Engine {
	r := gin.Default()

	// CORS middleware configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       // Allow all origins (use specific domains if needed)
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allowed HTTP methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Allowed headers
		ExposeHeaders:    []string{"X-Total-Count"},                           // Expose specific headers (optional)
		AllowCredentials: true,                                                // Allow credentials (cookies)
	}))

	// Routes

	r.GET("/getAllMeals", handler.GetAllMeals) // Get all meals for a user
	r.GET("/meal", handler.GetMeal)            // Get a single meal by ID for a user
	r.PUT("/meal", handler.UpdateMeal)         // Update a user's meal
	r.DELETE("/meal", handler.DeleteMeal)      //Delete a user's meal

	r.POST("/store-prediction", handler.StorePrediction) //successfully predicted images will be sent using this endpoint and this also stores meals
	return r
}
