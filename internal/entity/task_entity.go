package entity

import "time"

type TaskEntity struct {
	ID        int        `json:"id"`
	UserID    int        `json:"user_id"`
	Summary   string     `json:"summary"`
	IsDone    bool       `json:"is_done"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	UpdatedBy int        `json:"updated_by,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	DeletedBy *int       `json:"deleted_by,omitempty"`
}
