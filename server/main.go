package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DatabasePath = "./currency.db"
	ServerPort   = ":8080"
)

func main() {
	db, err := sql.Open("sqlite3", DatabasePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(CreateTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	currencyHandler := NewCurrencyHandler(db)

	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", currencyHandler.GetQuote)

	log.Printf("Server starting on %s\n", ServerPort)
	http.ListenAndServe(ServerPort, mux)
}
