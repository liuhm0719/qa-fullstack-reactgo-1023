// Backend reads ASTRA_BASE_PATH (injected by the Astra platform under
// path-based routing) and registers its routes under it.
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	base := strings.TrimRight(os.Getenv("ASTRA_BASE_PATH"), "/")
	mux := http.NewServeMux()
	mux.HandleFunc(base+"/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "hello from go backend"})
	})
	mux.HandleFunc(base+"/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})
	log.Fatal(http.ListenAndServe(":8080", mux))
}
