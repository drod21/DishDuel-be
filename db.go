package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/drod21/DishDuel-be/db"
	_ "github.com/lib/pq"
)

func dbInit() error {
	dbURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	var err error
	db.DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		return err
	}

	return db.DB.Ping()
}
