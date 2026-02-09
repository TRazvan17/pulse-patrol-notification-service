package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type CreateReq struct {
	To      string `json:"to"`
	Message string `json:"message"`
}

type CreateResp struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req CreateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	if req.To == "" || req.Message == "" {
		http.Error(w, "missing to/message", http.StatusBadRequest)
		return
	}

	resp := CreateResp{
		ID:        "notif-1",
		Status:    "queued",
		CreatedAt: time.Now().UTC(),
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", health)
	mux.HandleFunc("/notifications", create)

	log.Println("REST listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}