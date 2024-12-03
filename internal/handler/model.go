package handler

import "time"

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

// Mock data for demonstration purposes
var mockMeals = []Meal{
	{
		Id:        1,
		MealType:  "Breakfast",
		MealItems: []Food{{Id: 1, Name: "Eggs", Description: "Scrambled eggs", Calories: 200}},
		TimeEaten: time.Date(2024, 1, 1, 8, 0, 0, 0, time.UTC),
		CreatedAt: time.Date(2024, 1, 1, 7, 30, 0, 0, time.UTC),
	},
	{
		Id:        2,
		MealType:  "Dinner",
		MealItems: []Food{{Id: 3, Name: "Steak", Description: "Grilled steak", Calories: 900}, {Id: 2, Name: "Salad", Description: "Caesar salad", Calories: 150}},
		TimeEaten: time.Date(2024, 1, 2, 19, 0, 0, 0, time.UTC),
		CreatedAt: time.Date(2024, 1, 2, 18, 30, 0, 0, time.UTC),
	},
}
