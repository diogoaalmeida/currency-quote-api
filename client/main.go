package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

const (
	ServerURL     = "http://localhost:8080/cotacao"
	ClientTimeout = 300 * time.Millisecond
	OutputFile    = "cotacao.txt"
)

type BidResponse struct {
	Value string `json:"value"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), ClientTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", ServerURL, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	requestId := uuid.New().String()
	req.Header.Add("Request-Id", requestId)

	log.Printf("Making quotation request. Request-Id: %v\n", requestId)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	var bidResp BidResponse
	if err := json.Unmarshal(body, &bidResp); err != nil {
		log.Fatalf("Error parsing response: %v", err)
	}

	content := fmt.Sprintf("Dólar: %s", bidResp.Value)
	if err := os.WriteFile(OutputFile, []byte(content), 0644); err != nil {
		log.Fatalf("Error writing file: %v", err)
	}

	fmt.Printf("Saved to %s: %s\n", OutputFile, content)
}
