package repository

import (
	"context"
	"database/sql"
	"tasks-api/internal/entity"
	"tasks-api/internal/infra/database/queries"
)

type UserRepository struct {
	q *queries.Queries
}

func NewUserRepository(q *queries.Queries) *UserRepository {
	return &UserRepository{q: q}
}

func (u *UserRepository) NewUser(input entity.UserEntity) error {
	// I'm passing the ctx here because we aren't thinking about any traceability right now
	_, err := u.q.StoreUser(context.Background(), queries.StoreUserParams{
		FirstName: input.FirstName,
		Email:     sql.NullString{String: input.Email, Valid: true},
		Password:  input.Password,
		RoleID:    int32(input.RoleID),
	})

	return err
}

func (u *UserRepository) UserByEmail(email string) (*entity.UserEntity, error) {
	// I'm passing the ctx here because we aren't thinking about any traceability right now
	user, err := u.q.UserByEmail(context.Background(), sql.NullString{String: email, Valid: true})
	if err != nil {
		return nil, err
	}

	return &entity.UserEntity{
		ID:        int(user.ID),
		FirstName: user.FirstName,
		Email:     user.Email.String,
		Password:  user.Password,
		RoleID:    int(user.RoleID),
	}, nil
}
