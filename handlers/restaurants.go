package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/drod21/DishDuel-be/models"
)

func GetRestaurants(w http.ResponseWriter, r *http.Request) {
	// Implement database query to fetch restaurants
	restaurants := []models.Restaurant{
		// Fetch from database
	}
	json.NewEncoder(w).Encode(restaurants)
}
