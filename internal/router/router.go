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
	//Users
	// r.POST("/register", handler.Register) // Register a new user
	// r.POST("/login", handler.Login)       // Login a user


    //Meals
	r.POST("/getAllMeals", handler.GetAllMeals) // Get all meals for a user
	r.GET("/meal", handler.GetMeal)            // Get a single meal by ID for a user
	r.PUT("/meal", handler.UpdateMeal)         // Update a user's meal
	r.DELETE("/meal", handler.DeleteMeal)      //Delete a user's meal

	//can also be called -store-meal
	r.POST("/store-prediction", handler.StorePrediction) //successfully predicted images will be sent using this endpoint and this also stores meals

	r.GET("/getMacros/:mealName", handler.GetMealMacros) // Get the macros for a meal
	r.POST("/addMacros", handler.AddMealMacros)          // Add macros for a meal
	r.POST("/addBulkMacros", handler.AddBulkMealMacros) //Adds macros in bulk given a json of macros
	r.GET("/getAllMealMacros", handler.GetAllMealMacros) // Get all meal macros

	//Get Image using the image name
	r.GET("/getImage/:imageName", handler.GetImage)
	return r
}
