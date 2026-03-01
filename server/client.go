package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const CurrencyAPIURL = "https://economia.awesomeapi.com.br/json/last/USD-BRL"

func FetchCurrencyQuote(ctx context.Context) (*USDtoBRLResponse, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", CurrencyAPIURL, nil)
	if err != nil {
		log.Printf("Error creating request: %v\n", err)
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error making request: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response: %v\n", err)
		return nil, err
	}

	var quotationResponse USDtoBRLResponse
	err = json.Unmarshal(body, &quotationResponse)
	if err != nil {
		log.Printf("Error parsing json to struct: %v\n", err)
		return nil, err
	}
	return &quotationResponse, nil
}
