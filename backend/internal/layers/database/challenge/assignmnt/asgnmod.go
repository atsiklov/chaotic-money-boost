package assignment

import (
	enums "backend/internal/layers"
	"time"
)

type Assignment struct {
	ID         int64                  `db:"id"`
	UserID     int64                  `db:"user_id"`
	InstID     int64                  `db:"challenge_instance_id"`
	Status     enums.AssignmentStatus `db:"status"`
	CreatedAt  time.Time              `db:"created_at"`
	UpdatedAt  time.Time              `db:"updated_at"`
	Submission string                 `db:"submission"`
}
