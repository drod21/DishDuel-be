package main

import (
	"log"
	"net/http"

	"github.com/drod21/DishDuel-be/handlers"
	"github.com/gorilla/mux"
)

func startServer() error {
	router := mux.NewRouter()

	router.HandleFunc("/restaurants", handlers.GetRestaurants).Methods("GET")
	router.HandleFunc("/duel", handlers.DuelRestaurants).Methods("POST")

	log.Println("Server starting on :8080")
	return http.ListenAndServe(":8080", router)
}
