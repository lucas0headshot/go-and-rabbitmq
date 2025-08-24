package rabbitmq

import (
	"go-and-rabbitmq/src/util"

	"github.com/rabbitmq/amqp091-go"
)

func GetConnection() *amqp091.Connection {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	util.FailOnError(err, "Failed to connect to RabbitMQ")
	return conn
}
