package logger

import (
	"encoding/json"
	config "github.com/go-elk/core"
	"github.com/go-elk/core/logger/model"
	"github.com/streadway/amqp"
)

type RabbitMQClient struct{}

var connectRabbitMQ *amqp.Connection
var channelRabbitMQ *amqp.Channel

func NewLoggerRabbitMQ() *RabbitMQClient {

	return &RabbitMQClient{}
}

func (s *RabbitMQClient) Write(model model.Log) {
	amqpServerUrl := config.Config("RABBITMQ_URL")
	connectRabbitMQ, err = amqp.Dial(amqpServerUrl)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	channelRabbitMQ, err = connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	_, err = channelRabbitMQ.QueueDeclare(
		config.Config("RABBITMQ_QUEUE_NAME"),
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	msg, _ := json.Marshal(model)
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        msg,
	}

	if err := channelRabbitMQ.Publish(
		"",
		config.Config("RABBITMQ_QUEUE_NAME"),
		false,
		false,
		message,
	); err != nil {
	}
}
