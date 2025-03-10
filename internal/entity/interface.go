package entity

type UserRepository interface {
	NewUser(input UserEntity) error
	UserByEmail(email string) (*UserEntity, error)
}

type PaginationFilter struct {
	Page  int
	Limit int
}
type TechnicianRepository interface {
	NewTask(input TaskEntity) (int, error)
	FindTask(taskID, userID int) (*TaskEntity, error)
	CountTasksByUser(userID int) (int, error)
	UpdateTask(input TaskEntity) (int, error)
	AllTasksByUser(userID int, filter PaginationFilter) (*[]TaskEntity, error)
}

type ManagerRepository interface {
	DeleteTask(taskId, updatedBy int) error
	AllTasks(filter PaginationFilter) (*[]TaskEntity, error)
	CountTasks() (int, error)
}
