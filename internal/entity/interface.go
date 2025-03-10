package entity

type UserRepository interface {
	NewUser(input UserEntity) error
	UserByEmail(email string) (*UserEntity, error)
}
