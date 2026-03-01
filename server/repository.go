package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
)

const (
	CreateTableSQL = `
		CREATE TABLE IF NOT EXISTS currency_quotes (
			id TEXT NOT NULL PRIMARY KEY,
			from_currency TEXT,
			to_currency TEXT,
			currency_pair TEXT,
			bid_value TEXT,
			created_at TIMESTAMP,
			request_duration_ms INTEGER,
			data BLOB
		);
	`

	InsertQuoteSQL = `
		INSERT INTO currency_quotes (id, from_currency, to_currency, currency_pair, bid_value, request_duration_ms, data, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
)

func SaveCurrencyQuote(ctx context.Context, db *sql.DB, q *CurrencyQuote, requestDuration int64) error {
	jsonData, err := json.Marshal(q)
	if err != nil {
		log.Printf("Error parsing object to json: %v\n", err)
		return err
	}

	id := uuid.New().String()
	_, err = db.ExecContext(
		ctx,
		InsertQuoteSQL,
		id,
		q.Code,
		q.Codein,
		q.Name,
		q.Bid,
		requestDuration,
		jsonData,
		time.Now(),
	)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return err
	}
	return nil
}
