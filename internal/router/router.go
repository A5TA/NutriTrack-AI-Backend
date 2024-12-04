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
	r.POST("/meal", handler.PostMeal)
	//Handle both the single and multiple day requests with this handler
	r.GET("/meals/:startDate", handler.GetAllMeals)
	r.GET("/meals/:startDate/:endDate", handler.GetAllMeals)

	r.GET("/meal/:meal_id", handler.GetMeal)
	r.PUT("/meal/:meal_id", handler.UpdateMeal) // Use PUT for updates
	r.DELETE("/meal/:meal_id", handler.DeleteMeal)

	//Ai predictions
	r.POST("/store-prediction", handler.PredictFood)
	return r
}
