package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const (
	APITimeout = 200 * time.Millisecond
	DBTimeout  = 10 * time.Millisecond
)

type CurrencyHandler struct {
	db *sql.DB
}

func NewCurrencyHandler(db *sql.DB) *CurrencyHandler {
	return &CurrencyHandler{db: db}
}

func (h *CurrencyHandler) GetQuote(w http.ResponseWriter, r *http.Request) {
	ctxAPI, cancelAPI := context.WithTimeout(context.Background(), APITimeout)
	defer cancelAPI()

	requestId := r.Header.Get("Request-Id")
	log.Printf("Receiving quotation request. Request-Id: %v\n", requestId)

	start := time.Now()
	quotation, err := FetchCurrencyQuote(ctxAPI)
	duration := time.Since(start)
	if err != nil {
		log.Printf("Error getting quotation: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(APIError{Message: "Something went wrong."})
		return
	}

	ctxDB, cancelDB := context.WithTimeout(context.Background(), DBTimeout)
	defer cancelDB()

	err = SaveCurrencyQuote(ctxDB, h.db, &quotation.UsdBrl, duration.Milliseconds())
	if err != nil {
		log.Printf("Error saving quotation: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(APIError{Message: "Something went wrong."})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(BidResponse{Value: quotation.UsdBrl.Bid})
}
