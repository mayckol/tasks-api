package notify

import "fmt"

type SimpleNotifier struct {
}

func (s SimpleNotifier) TaskPerformed(taskID, userID int) error {
	fmt.Printf("task %d performed by user %d\n", taskID, userID)
	return nil
}
