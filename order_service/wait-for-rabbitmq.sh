#!/bin/sh

# Wait for RabbitMQ to be available and healthy
until nc -z -v -w30 rabbitmq 5672 >/dev/null 2>&1; do
  echo "Waiting for RabbitMQ to be available..."
  sleep 1
done

echo "RabbitMQ is available"

# Start the order service
exec "$@"
