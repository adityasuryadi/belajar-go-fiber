package config

import (
	"go-blog/exception"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewRabbitmqConn(configuration Config) *amqp.Connection {
	host := configuration.Get("RABBITMQ_HOST")
	user := configuration.Get("RABBITMQ_USER")
	pass := configuration.Get("RABBITMQ_PASS")
	port := configuration.Get("RABBITMQ_PORT")

	conn, err := amqp.Dial("amqp://" + user + ":" + pass + "@" + host + ":" + port)
	exception.PanicIfNeeded(err)

	return conn
}
