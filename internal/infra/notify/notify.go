package notify

import (
	"fmt"
	"tasks-api/internal/infra/messaging"
)

type SimpleNotifier struct {
	messaging messaging.Messaging
}

func NewSimpleNotifier(messaging messaging.Messaging) SimpleNotifier {
	return SimpleNotifier{messaging: messaging}
}

func (s SimpleNotifier) TaskPerformed(taskID, userID int) error {
	// for testing purposes, we are not going to send the message, just print it, otherwise we would need to mock the messaging service
	if s.messaging == nil {
		return nil
	}
	return s.messaging.Send([]byte(fmt.Sprintf("task %d performed by user %d\n", taskID, userID)), "tasks")
}
