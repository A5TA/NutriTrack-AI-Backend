package router

import (
	"github.com/A5TA/NutriTrack-Ai-backend/internal/handler"
	"github.com/gin-gonic/gin"
)

// New creates a new Gin router and sets up routes
func New() *gin.Engine {
	r := gin.Default()

	// Routes
	r.POST("/meal", handler.PostMeal)
	//Handle both the single and multiple day requests with this handler
	r.GET("/meals/:startDate", handler.GetAllMeals)
	r.GET("/meals/:startDate/:endDate", handler.GetAllMeals)
	
	r.GET("/meal/:meal_id", handler.GetMeal)
	r.PUT("/meal/:meal_id", handler.UpdateMeal) // Use PUT for updates
	r.DELETE("/meal/:meal_id", handler.DeleteMeal)

	return r
}
