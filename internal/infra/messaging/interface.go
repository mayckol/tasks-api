package messaging

import "tasks-api/internal/infra/messaging/rabbitmqpkg"

type Messaging interface {
	Send(content []byte, queueName string) error
	Consume(queueName string) (<-chan rabbitmqpkg.Message, error)
	Cancel(consumerTag string) error
}
