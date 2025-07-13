package user

import "time"

type User struct {
	ID        int64     `db:"id"`
	Nickname  string    `db:"nickname"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
}
