package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/A5TA/NutriTrack-Ai-backend/internal/handler"
)


type test_case struct {
	name string
	status int
}

func Test_PostMeal(t *testing.T) {
	testCases := []test_case{
		{
			name: "not implemented",
			status: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//Arrange
			r := httptest.NewRequest(http.MethodPost, "/", nil)
			w := httptest.NewRecorder()
			//Act
			handler.PostMeal()(w,r)

			//Assert
			if w.Result().StatusCode != tc.status {
				t.Errorf("Expected: %d, but got: %d", tc.status, w.Result().StatusCode)
			}
		})
	}
}

func Test_GetAllMeals(t *testing.T) {
	testCases := []test_case{
		{
			name: "not implemented",
			status: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//Arrange
			r := httptest.NewRequest(http.MethodPost, "/", nil)
			w := httptest.NewRecorder()
			//Act
			handler.GetAllMeals()(w,r)

			//Assert
			if w.Result().StatusCode != tc.status {
				t.Errorf("Expected: %d, but got: %d", tc.status, w.Result().StatusCode)
			}
		})
	}


}
func Test_GetMeal(t *testing.T) {
	testCases := []test_case{
		{
			name: "not implemented",
			status: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//Arrange
			r := httptest.NewRequest(http.MethodPost, "/", nil)
			w := httptest.NewRecorder()
			//Act
			handler.GetMeal()(w,r)

			//Assert
			if w.Result().StatusCode != tc.status {
				t.Errorf("Expected: %d, but got: %d", tc.status, w.Result().StatusCode)
			}
		})
	}


}
func Test_UpdateMeal(t *testing.T) {
	testCases := []test_case{
		{
			name: "not implemented",
			status: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//Arrange
			r := httptest.NewRequest(http.MethodPost, "/", nil)
			w := httptest.NewRecorder()
			//Act
			handler.UpdateMeal()(w,r)

			//Assert
			if w.Result().StatusCode != tc.status {
				t.Errorf("Expected: %d, but got: %d", tc.status, w.Result().StatusCode)
			}
		})
	}


}
func Test_DeleteMeal(t *testing.T) {
	testCases := []test_case{
		{
			name: "not implemented",
			status: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//Arrange
			r := httptest.NewRequest(http.MethodPost, "/", nil)
			w := httptest.NewRecorder()
			//Act
			handler.DeleteMeal()(w,r)

			//Assert
			if w.Result().StatusCode != tc.status {
				t.Errorf("Expected: %d, but got: %d", tc.status, w.Result().StatusCode)
			}
		})
	}


}

