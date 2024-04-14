package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func placeOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	order, err := decodeOrderRequest(r)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = publishOrderToRabbitMQ(order)
	if err != nil {
		http.Error(w, "Failed to publish order", http.StatusInternalServerError)
		return
	}

	log.Printf("Order received: %s", order.Food)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
