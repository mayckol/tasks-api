package rabbitmqpkg

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"tasks-api/configs"
)

type Connection struct {
	envs *configs.EnvVars
}

func NewConnection(envs *configs.EnvVars) *Connection {
	return &Connection{envs: envs}
}

func (c *Connection) Dial() *amqp.Connection {
	conn, err := amqp.Dial(
		fmt.Sprintf("amqp://%s:%s@%s:%s/",
			c.envs.RabbitmqDefaultUser,
			c.envs.RabbitmqDefaultPass,
			c.envs.RabbitmqDefaultHost,
			c.envs.RabbitmqDefaultPort,
		),
	)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to RabbitMQ: %s", err))
	}

	return conn
}
