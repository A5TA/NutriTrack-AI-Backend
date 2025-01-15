package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/A5TA/NutriTrack-Ai-backend/internal/handler"
	"github.com/gin-gonic/gin"
)

type Meal struct {
	Id        int       `json:"id"`
	MealType  string    `json:"mealType"`      //type of meal (lunch, dinner, breakfast)
	MealItems []Food    `json:"mealItems"` // Array of food items in the meal
	TimeEaten time.Time `json:"timeEaten"` //time eaten
	CreatedAt time.Time `json:"createdAt"`
}

type Food struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`        // Name of the food item
	Description string `json:"description"` // Optional description of the food if wanted to add for future enhancements
	Calories    int    `json:"calories"`    // Caloric value
}


func Test_PostMeal(t *testing.T) {
	type testCase struct {
		name       string
		input      Meal
		status     int
		expectBody bool
	}
	testCases := []testCase{
		{
			name: "valid meal",
			input: Meal{
				Id:        1,
				MealType:  "lunch",
				MealItems: []Food{{Name: "Chicken", Calories: 200}},
				TimeEaten: time.Now(),
				CreatedAt: time.Now(),
			},
			status:     http.StatusOK,
			expectBody: true,
		},
		{
			name:       "invalid meal (empty body)",
			input:      Meal{}, // Empty body will cause BindJSON to fail
			status:     http.StatusBadRequest,
			expectBody: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			r := gin.Default()
			r.POST("/", handler.PostMeal)

			var reqBody []byte
			var err error
			if tc.expectBody {
				reqBody, err = json.Marshal(tc.input)
				if err != nil {
					t.Fatalf("Failed to marshal input: %v", err)
				}
			}

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			// Act
			r.ServeHTTP(w, req)

			// Assert
			if w.Code != tc.status {
				t.Errorf("Expected: %d, but got: %d", tc.status, w.Code)
			}

			if tc.expectBody {
				var responseBody Meal
				if err := json.Unmarshal(w.Body.Bytes(), &responseBody); err != nil {
					t.Fatalf("Failed to unmarshal response body: %v", err)
				}

				if responseBody.MealType != tc.input.MealType {
					t.Errorf("Expected mealType: %s, but got: %s", tc.input.MealType, responseBody.MealType)
				}
			}
		})
	}
}

func Test_GetAllMeals(t *testing.T) {
	type testCase struct {
		name         string
		initialMeals []Meal
		startDate    string
		endDate      string
		status       int
		expectEmpty  bool
	}

	// Mock meals data to simulate different meal entries
	mockMeals := []Meal{
		{
			Id:        1,
			MealType:  "breakfast",
			MealItems: []Food{{Name: "Eggs", Calories: 155}},
			TimeEaten: time.Date(2024, time.December, 16, 8, 0, 0, 0, time.UTC),
			CreatedAt: time.Now(),
		},
		{
			Id:        2,
			MealType:  "lunch",
			MealItems: []Food{{Name: "Sandwich", Calories: 250}},
			TimeEaten: time.Date(2024, time.December, 17, 12, 0, 0, 0, time.UTC),
			CreatedAt: time.Now(),
		},
		{
			Id:        3,
			MealType:  "dinner",
			MealItems: []Food{{Name: "Steak", Calories: 500}},
			TimeEaten: time.Date(2024, time.December, 18, 19, 0, 0, 0, time.UTC),
			CreatedAt: time.Now(),
		},
	}

	testCases := []testCase{
		{
			name:         "no meals available",
			initialMeals: []Meal{}, // No meals should match
			startDate:    "2024-12-16", // Sample date
			endDate:      "2024-12-16", // Same date, no meals in this range
			status:       http.StatusOK,
			expectEmpty:  true,
		},
		{
			name:         "meals available within range",
			initialMeals: mockMeals,
			startDate:    "2024-12-16",
			endDate:      "2024-12-17", // One meal should match (breakfast)
			status:       http.StatusOK,
			expectEmpty:  false,
		},
		{
			name:         "meals available for a date range",
			initialMeals: mockMeals,
			startDate:    "2024-12-16",
			endDate:      "2024-12-18", // Two meals should match (breakfast and lunch)
			status:       http.StatusOK,
			expectEmpty:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			r := gin.Default()
			r.GET("/:startDate/:endDate", handler.GetAllMeals)

			// Create the request with the URL parameters
			url := "/" + tc.startDate + "/" + tc.endDate
			req := httptest.NewRequest(http.MethodGet, url, nil)
			w := httptest.NewRecorder()

			// Act
			r.ServeHTTP(w, req)

			// Assert the status code
			if w.Code != tc.status {
				t.Errorf("Expected status: %d, but got: %d", tc.status, w.Code)
			}

			// Unmarshal the response body
			var responseBody []Meal
			if err := json.Unmarshal(w.Body.Bytes(), &responseBody); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}
			

			// Assert that the number of meals matches expectation
			if tc.expectEmpty && len(responseBody) != 0 {
				t.Errorf("Expected no meals, but got: %v", responseBody)
			}
		})
	}
}

// Similar test cases for `GetMeal`, `UpdateMeal`, and `DeleteMeal`
