package service

type RabbitMqService interface {
	PublishQueue(queueName string, data interface{}) error
}
