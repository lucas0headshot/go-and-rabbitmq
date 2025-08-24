package rabbitmq

import (
	"go-and-rabbitmq/src/util"

	"github.com/rabbitmq/amqp091-go"
)

func OpenChannel(conn *amqp091.Connection) *amqp091.Channel {
	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open a channel")
	return ch
}
