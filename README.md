# currency-quote-api

A Go client-server application that fetches and persists USD-BRL exchange rates.

## Architecture

- **server/**: HTTP server that fetches the current USD-BRL quote from an external API and persists it to SQLite
- **client/**: HTTP client that requests the quote from the server and saves it to a local file

## Requirements

- Go 1.21+

## Running

### Start the server

```bash
cd server
go run *.go
```

The server starts on port `8080` and exposes a single endpoint:

```
GET http://localhost:8080/cotacao
```

Response:
```json
{"value": "5.8950"}
```

### Run the client

In a separate terminal:

```bash
cd client
go run main.go
```

This creates `client/cotacao.txt` with content like:

```
Dólar: 5.8950
```

## Timeouts

| Operation | Timeout |
|---|---|
| Server → external API | 200ms |
| Server → SQLite write | 10ms |
| Client → server | 300ms |

All timeouts log an error if exceeded.
