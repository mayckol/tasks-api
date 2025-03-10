package main

import (
	"context"
	"database/sql"
	"fmt"
	"tasks-api/configs"
	"tasks-api/internal/infra/database"
	"tasks-api/internal/infra/database/queries"
	"tasks-api/internal/role"
	"tasks-api/utils"
)

func main() {
	envs := configs.LoadEnv()

	db := database.New(envs)
	defer db.Close()

	q := queries.New(db)
	q.StoreRole(context.Background(), queries.StoreRoleParams{
		ID:    role.MANAGER,
		Alias: "MANAGER",
	})
	q.StoreRole(context.Background(), queries.StoreRoleParams{
		ID:    role.TECHNICIAN,
		Alias: "TECHNICIAN",
	})

	pass, _ := utils.HashPassword("managersecret")
	q.StoreUser(context.Background(), queries.StoreUserParams{
		FirstName: "Manager",
		Email: sql.NullString{
			String: "manager@test.com",
			Valid:  true,
		},
		Password: pass,
		RoleID:   role.MANAGER,
	})

	pass, _ = utils.HashPassword("techniciansecret")
	q.StoreUser(context.Background(), queries.StoreUserParams{
		FirstName: "Technician",
		Email: sql.NullString{
			String: "technician@test.com",
			Valid:  true,
		},
		Password: pass,
		RoleID:   role.TECHNICIAN,
	})

	fmt.Println("seeded roles and users")
}
