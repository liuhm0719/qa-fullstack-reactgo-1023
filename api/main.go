// Backend reads ASTRA_BASE_PATH (injected by the Astra platform under
// path-based routing) and registers its routes under it.
package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	base := strings.TrimRight(os.Getenv("ASTRA_BASE_PATH"), "/")
	mux := http.NewServeMux()

	mux.HandleFunc(base+"/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "hello from go backend"})
	})

	mux.HandleFunc(base+"/api/db/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		dsn := os.Getenv("DATABASE_URL")
		if dsn == "" {
			w.WriteHeader(http.StatusServiceUnavailable)
			json.NewEncoder(w).Encode(map[string]string{"error": "DATABASE_URL not set"})
			return
		}
		db, err := sql.Open("postgres", dsn)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		defer db.Close()

		var dbName string
		var serverTime time.Time
		err = db.QueryRow("SELECT current_database(), now()").Scan(&dbName, &serverTime)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		json.NewEncoder(w).Encode(map[string]string{
			"database":    dbName,
			"server_time": serverTime.UTC().Format(time.RFC3339),
		})
	})

	mux.HandleFunc(base+"/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
