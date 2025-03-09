package usecase

import (
	"net/http"
	"tasks-api/internal/auth/jwtpkg"
	"tasks-api/internal/entity"
	"tasks-api/internal/errorpkg"
	"tasks-api/utils"
	"time"
)

// this needs to be in a config file (it's just an example)
var duration = time.Duration(24) * time.Hour

// swagger:model SigninInputDTO
type SigninInputDTO struct {
	Email    string `validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=50"`
}

type SigninOutputDTO struct {
	AccessToken string `json:"access_token"`
	RoleID      int    `json:"role_id"`
}

type SigninUseCase struct {
	UserRepository entity.UserRepository
	JWTService     jwtpkg.TokenServiceInterface
}

func (s *SigninUseCase) Execute(input SigninInputDTO) (*SigninOutputDTO, *errorpkg.AppError) {
	var user *entity.UserEntity
	if input.Email != "" {
		u, err := s.UserRepository.UserByEmail(input.Email)
		if err != nil {
			return nil, errorpkg.New("something went wrong", http.StatusInternalServerError, err)
		}
		user = u
	}

	if user == nil {
		return nil, errorpkg.New("user not found", http.StatusNotFound, nil)
	}

	if err := utils.CheckPasswordHash(input.Password, user.Password); err != nil {
		return nil, errorpkg.New("invalid password", http.StatusUnauthorized, nil)
	}

	userClaims, err := s.userClaims(*user, user.RoleID)
	if err != nil {
		return nil, errorpkg.New("something went wrong", http.StatusInternalServerError, err)
	}

	token, accessClaims, err := s.JWTService.GenerateToken(*userClaims)
	if err != nil {
		return nil, nil
	}

	return &SigninOutputDTO{
		AccessToken: token,
		RoleID:      accessClaims.RoleID,
	}, nil
}

func (s *SigninUseCase) userClaims(entityUser entity.UserEntity, roleID int) (*jwtpkg.UserClaims, error) {
	userClaims, err := jwtpkg.NewUserClaims(
		entityUser.ID,
		roleID,
		duration,
	)
	if err != nil {
		return nil, err
	}
	return userClaims, nil
}
