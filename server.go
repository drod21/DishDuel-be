package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"math"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Restaurant struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	MMR  int    `json:"mmr"`
}

type DuelRequest struct {
	WinnerID string `json:"winner_id"`
	LoserID  string `json:"loser_id"`
}

var db *sql.DB

func main() {
	// Initialize database connection
	var err error
	db, err = sql.Open("postgres", "postgres://username:password@localhost/dishduel?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/restaurants", getRestaurants).Methods("GET")
	router.HandleFunc("/duel", duelRestaurants).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getRestaurants(w http.ResponseWriter, r *http.Request) {
	// Implement database query to fetch restaurants
}

func duelRestaurants(w http.ResponseWriter, r *http.Request) {
	var duelReq DuelRequest
	err := json.NewDecoder(r.Body).Decode(&duelReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	winner, loser := findRestaurants(duelReq.WinnerID, duelReq.LoserID)
	if winner == nil || loser == nil {
		http.Error(w, "Invalid restaurant IDs", http.StatusBadRequest)
		return
	}

	updateMMR(winner, loser)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"winner": winner,
		"loser":  loser,
	})
}

func findRestaurants(winnerID, loserID string) (*Restaurant, *Restaurant) {
	var winner, loser *Restaurant
	// Implement database query to find restaurants
	return winner, loser
}

func updateMMR(winner, loser *Restaurant) {
	kFactor := 32.0
	expectedScoreWinner := 1 / (1 + math.Pow(10, float64(loser.MMR-winner.MMR)/400))
	expectedScoreLoser := 1 - expectedScoreWinner

	winner.MMR += int(kFactor * (1 - expectedScoreWinner))
	loser.MMR += int(kFactor * (0 - expectedScoreLoser))

	// Implement database transactions for updating MMR and recording duels
	// Start a database transaction
	tx, err := db.Begin()
	if err != nil {
		log.Printf("Error starting transaction: %v", err)
		return
	}
	defer tx.Rollback() // Rollback the transaction if it hasn't been committed

	// Update winner's MMR
	_, err = tx.Exec("UPDATE restaurants SET mmr = ? WHERE id = ?", winner.MMR, winner.ID)
	if err != nil {
		log.Printf("Error updating winner's MMR: %v", err)
		return
	}

	// Update loser's MMR
	_, err = tx.Exec("UPDATE restaurants SET mmr = ? WHERE id = ?", loser.MMR, loser.ID)
	if err != nil {
		log.Printf("Error updating loser's MMR: %v", err)
		return
	}

	// Record the duel
	_, err = tx.Exec("INSERT INTO duels (winner_id, loser_id, winner_mmr, loser_mmr) VALUES (?, ?, ?, ?)",
		winner.ID, loser.ID, winner.MMR, loser.MMR)
	if err != nil {
		log.Printf("Error recording duel: %v", err)
		return
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Printf("Error committing transaction: %v", err)
		return
	}

	log.Printf("Successfully updated MMR and recorded duel for winner %s and loser %s", winner.ID, loser.ID)
}
