package user

import (
	myerr "backend/internal/errors"
	myconst "backend/internal/layers/database"
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5/pgconn"

	"github.com/jackc/pgx/v5"
)

type pgRepository struct {
	conn *pgx.Conn
}

func (repo *pgRepository) Create(ctx context.Context, user *User) error {
	nick := user.Nickname
	mail := user.Email

	query := `
		insert into users (nickname, email) values ($1, $2)
		returning id, created_at
	`
	err := repo.conn.QueryRow(ctx, query, nick, mail).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == myconst.UNIQUE_VIOLATION_CODE {
			log.Printf("Failed to create user with nickname=%s: unique constraint violation", nick)
			return myerr.ErrUserAlreadyExists
		}
		log.Printf("Failed to create user with nickname=%s: %s", nick, err.Error())
		return err
	}
	log.Printf("Created new user with nickname=%s", nick)
	return nil
}

func (repo *pgRepository) FindByID(ctx context.Context, id int64) (*User, error) {
	query := `select id, nickname, email, created_at from users where id = $1`
	row := repo.conn.QueryRow(ctx, query, id)

	var user User
	err := row.Scan(&user.ID, &user.Nickname, &user.Email, &user.CreatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Printf("No user with id=%d", id)
			return nil, myerr.ErrUserNotFound
		}
		log.Printf("Failed to find user with id=%d: %s", id, err.Error())
		return nil, err
	}
	return &user, nil
}

func NewPgRepo(conn *pgx.Conn) Repository {
	return &pgRepository{conn: conn}
}
