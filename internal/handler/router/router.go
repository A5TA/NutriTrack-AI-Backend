package router

import (
	"net/http"

	"github.com/A5TA/NutriTrack-Ai-backend/internal/handler"
)

//Creates a new Router with the handlers all configured
func New() *http.ServeMux {
	mux := http.NewServeMux()

	//Create Routes
	mux.HandleFunc("POST /meal", handler.PostMeal())
	mux.HandleFunc("GET /meals/{startDate}/{endDate}", handler.GetAllMeals())
	mux.HandleFunc("GET /meal/{meal_id}", handler.GetMeal())
	mux.HandleFunc("UPDATE /meal/{meal_id}", handler.UpdateMeal())
	mux.HandleFunc("DELETE /meal/{meal_id}", handler.DeleteMeal())

	return mux
}