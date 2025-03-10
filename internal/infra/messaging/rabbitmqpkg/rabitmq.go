package rabbitmqpkg

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type Messaging struct {
	Client *amqp.Connection
}

func NewMessaging(client *amqp.Connection) *Messaging {
	return &Messaging{Client: client}
}

func (r *Messaging) Send(content []byte, queueName string) error {
	ch, err := r.Client.Channel()
	if err != nil {
		return err
	}
	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			// TODO it can better make this dynamic to accept any content type
			ContentType: "text/plain",
			Body:        content,
		})

	if err != nil {
		log.Printf(" [x] Sent %s\n", string(content))
	}

	return nil
}

type Message interface {
	Body() []byte
}

type deliveryAdapter struct {
	delivery amqp.Delivery
}

func (d *deliveryAdapter) Body() []byte {
	return d.delivery.Body
}

func (r *Messaging) Consume(queueName string) (<-chan Message, error) {
	ch, err := r.Client.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
	}

	out := make(chan Message)
	go func() {
		defer close(out)
		for msg := range msgs {
			out <- &deliveryAdapter{delivery: msg}
		}
	}()
	return out, nil
}

func (r *Messaging) Cancel(consumerTag string) error {
	ch, err := r.Client.Channel()
	if err != nil {
		return err
	}
	return ch.Cancel(consumerTag, false)
}
