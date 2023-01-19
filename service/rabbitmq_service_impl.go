package service

import (
	"go-blog/exception"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewRabbitMqService(conn *amqp.Connection) RabbitMqService {
	return &RabbitMqServiceImpl{
		connection: conn,
	}
}

type RabbitMqServiceImpl struct {
	connection *amqp.Connection
}

// PublishQueue implements RabbitMqService
func (service *RabbitMqServiceImpl) PublishQueue(queueName string, data interface{}) error {
	channel, err := service.connection.Channel()
	exception.PanicIfNeeded(err)
	defer channel.Close()

	_, err = channel.QueueDeclare(
		"TestQueue", //nama
		false,       //durable
		false,       //delete when unused
		false,       //exclusive
		false,       // no-wait
		nil,         //arguments
	)

	exception.PanicIfNeeded(err)

	// var dta interface{}
	// json.Unmarshal([]byte(dta.(string)), &data)
	err = channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(queueName),
		},
	)

	exception.PanicIfNeeded(err)

	return nil

}
