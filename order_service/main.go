package main

import (
	"log"
	"net/http"
)

const (
	QueueName = "orders"
	BrokerURL = "amqp://guest:guest@rabbitmq:5672/"
)

type Order struct {
	Food string `json:"food"`
}

func main() {
	err := declareQueue()
	if err != nil {
		log.Fatalf("Failed to declare the queue: %v", err)
	}

	http.HandleFunc("/order", placeOrderHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
