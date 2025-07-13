package instance

import (
	enums "backend/internal/layers"
	"time"
)

type ChallengeInstance struct {
	ID         int64                `db:"id"`
	TemplateID int64                `db:"challenge_template_id"`
	Status     enums.ChgeInstStatus `db:"status"`
	CreatedAt  time.Time            `db:"created_at"`
	UpdatedAt  time.Time            `db:"updated_at"`
	StartedAt  *time.Time           `db:"started_at"`
	ExpiresAt  *time.Time           `db:"expires_at"`
}
