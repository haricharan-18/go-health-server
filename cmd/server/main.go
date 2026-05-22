package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type HealthResponse struct {
	Status string `json:"status"`
	Node   string `json:"node"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	node := os.Getenv("NODE_ID")

	response := HealthResponse{
		Status: "ok",
		Node:   node,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/health", healthHandler)

	log.Println("Server running on port", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
