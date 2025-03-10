package main

import (
	"log"
	"tasks-api/configs"
	"tasks-api/internal/infra/messaging/rabbitmqpkg"
)

func main() {
	envs := configs.LoadEnv()

	messagingConnection := rabbitmqpkg.NewConnection(envs).Dial()
	defer messagingConnection.Close()

	messaging := rabbitmqpkg.NewMessaging(messagingConnection)
	log.Println("waiting for messages")
	msgs, err := messaging.Consume("tasks")
	if err != nil {
		log.Fatalf("failed to consume messages: %s", err)
	}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body())
		}
	}()
}
