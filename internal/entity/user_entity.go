package entity

type UserEntity struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	RoleID    int    `json:"role_id"`
}
