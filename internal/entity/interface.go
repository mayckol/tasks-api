package entity

type UserRepository interface {
	NewUser(input UserEntity) error
	UserByEmail(email string) (*UserEntity, error)
}

type TechnicianRepository interface {
	NewTask(input TaskEntity) (int, error)
	FindTask(taskID, userID int) (*TaskEntity, error)
	CountTasksByUser(userID int) (int, error)
	UpdateTask(input TaskEntity) (int, error)
	AllTasksByUser(userID, page int) (*[]TaskEntity, error)
}
