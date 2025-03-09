package usecase

import (
	"net/http"
	"tasks-api/internal/entity"
	"tasks-api/internal/errorpkg"
	"tasks-api/utils"
)

// swagger:model NewUserInputDTO

type NewUserInputDTO struct {
	FirstName string `json:"first_name" validate:"required,min=2,max=50"`
	// Here we'd need to validate "unique" in database
	Email string `validate:"required,email"`
	// Here we'd need to validate password strength
	Password string `json:"password" validate:"required,min=8,max=50"`
	// Here we'd need to validate that the role exists in the database
	RoleID int `json:"role_id" validate:"required"`
}

type NewUserUseCase struct {
	UserRepository entity.UserRepository
}

func (n *NewUserUseCase) Execute(input NewUserInputDTO) *errorpkg.AppError {
	hashPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return errorpkg.Wrap("failed to hash password", http.StatusInternalServerError, err)
	}
	input.Password = hashPassword
	err = n.UserRepository.NewUser(entity.UserEntity{
		FirstName: input.FirstName,
		Email:     input.Email,
		Password:  input.Password,
		RoleID:    input.RoleID,
	})

	if err != nil {
		return errorpkg.Wrap("failed to create new user", http.StatusInternalServerError, err)
	}

	return nil
}
