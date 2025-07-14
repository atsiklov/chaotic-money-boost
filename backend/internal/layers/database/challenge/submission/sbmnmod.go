package submission

import "time"

type Submission struct {
	ID          int64     `db:"id"`
	UserID      int64     `db:"user_id"`
	SubmittedAt time.Time `db:"submitted_at"`
	Submission  string    `db:"submittion"`
}
