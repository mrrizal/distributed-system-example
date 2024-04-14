package main

import (
	"encoding/json"
	"net/http"

	"github.com/streadway/amqp"
)

func decodeOrderRequest(r *http.Request) (Order, error) {
	var order Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		return Order{}, err
	}
	return order, nil
}

func declareQueue() error {
	conn, err := amqp.Dial(BrokerURL)
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

func publishOrderToRabbitMQ(order Order) error {
	conn, err := amqp.Dial(BrokerURL)
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	body, err := json.Marshal(order)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",
		QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return err
	}

	return nil
}
