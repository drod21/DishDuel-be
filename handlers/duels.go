package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/drod21/DishDuel-be/models"
	"github.com/drod21/DishDuel-be/utils"
)

type DuelRequest struct {
	WinnerID string `json:"winner_id"`
	LoserID  string `json:"loser_id"`
}

func DuelRestaurants(w http.ResponseWriter, r *http.Request) {
	var duelReq DuelRequest
	if err := json.NewDecoder(r.Body).Decode(&duelReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Fetch restaurants from database
	winner, loser := findRestaurants(duelReq.WinnerID, duelReq.LoserID)
	if winner == nil || loser == nil {
		http.Error(w, "Invalid restaurant IDs", http.StatusBadRequest)
		return
	}

	utils.UpdateMMR(winner, loser)

	// Update restaurants in database

	json.NewEncoder(w).Encode(map[string]interface{}{
		"winner": winner,
		"loser":  loser,
	})
}

func findRestaurants(winnerID, loserID string) (*models.Restaurant, *models.Restaurant) {
	// Implement database query to find restaurants
	return nil, nil
}
