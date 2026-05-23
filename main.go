package main

import (
	"fmt"
	"log"
	"net/http"

	"sei-ratelimiter/internal/config"
)

func main() {
	cfg := config.Load()

	fmt.Printf("starting sei-ratelimiter node=%s port=%s\n", cfg.NodeID, cfg.Port)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status":"ok","node":%q}`, cfg.NodeID)
	})

	addr := ":" + cfg.Port
	log.Printf("listening on %s", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
