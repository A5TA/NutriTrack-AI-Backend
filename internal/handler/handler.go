package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PostMeal(c *gin.Context) {
	var newMeal Meal

	if err := c.BindJSON(&newMeal); err != nil {
		return
	}

	mockMeals = append(mockMeals, newMeal)
	c.JSON(http.StatusOK, newMeal)
}

// GetAllMeals fetches meals between startDate and endDate
func GetAllMeals(c *gin.Context) {
	// Extract startDate and endDate from URL parameters
	startDateStr := c.Param("startDate")
	endDateStr := c.Param("endDate")

	// Parse dates
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid startDate format, expected YYYY-MM-DD"})
		return
	}
	// Parse endDate, or default it to startDate if not provided
	var endDate time.Time
	if endDateStr == "" {
		endDate = startDate
	} else {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid endDate format, expected YYYY-MM-DD"})
			return
		}
	}

	// Filter meals by date range
	var filteredMeals []Meal
	for _, meal := range mockMeals {
		if meal.TimeEaten.After(startDate) && meal.TimeEaten.Before(endDate.Add(24*time.Hour)) {
			filteredMeals = append(filteredMeals, meal)
		}
	}

	// Return the filtered meals
	c.JSON(http.StatusOK, filteredMeals)
}

func GetMeal(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

func UpdateMeal(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

func DeleteMeal(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}
