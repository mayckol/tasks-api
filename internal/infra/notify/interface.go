package notify

type NotifyInterface interface {
	TaskPerformed(taskID, userID int) error
}
