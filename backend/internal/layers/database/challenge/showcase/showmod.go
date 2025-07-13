package showcase

import "time"

type ShowcaseChallenge struct {
	InstanceID  int64
	Category    string
	Description string
	StartedAt   *time.Time
	ExpiresAt   *time.Time
}
