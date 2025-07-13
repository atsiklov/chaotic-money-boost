package template

import "time"

type ChallengeTemplate struct {
	ID          int64          `db:"id"`
	Category    string         `db:"category"`
	Description string         `db:"description"`
	Duration    *time.Duration `db:"duration"`
}
