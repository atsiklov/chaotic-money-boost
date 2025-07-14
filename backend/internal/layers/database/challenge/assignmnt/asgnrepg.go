package assignment

import (
	myerr "backend/internal/errors"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type pgRepository struct {
	conn *pgx.Conn
}

func (repo *pgRepository) Create(ctx context.Context, asignNew *Assignment) error {
	const query = `insert into challenge_assignments (challenge_instance_id, user_id, status) values ($1, $2, $3)`

	_, err := repo.conn.Exec(ctx, query, asignNew.InstID, asignNew.UserID, asignNew.Status)
	if err != nil {
		log.Printf("Failed to create new assignment: %s", err.Error())
		return err
	}
	return nil
}

func (repo *pgRepository) Update(ctx context.Context, asignUpd *Assignment) error {
	const query = `update challenge_assignments 
		set status = $1, submission = $2, submitted_at = now(), updated_at = now() 
		where challenge_instance_id = $3 and user_id = $4;
	`
	cmdTag, err := repo.conn.Exec(ctx, query, asignUpd.Status, asignUpd.Submission, asignUpd.InstID, asignUpd.UserID)
	if err != nil {
		log.Printf("Failed to update an assignment: %s", err.Error())
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		log.Printf("Failed to update: assignment doesn't exist for user id = %d and instance id = %d", asignUpd.UserID, asignUpd.InstID)
		return myerr.ErrNoRecordsToUpdate
	}
	return nil
}

func NewPgRepo(conn *pgx.Conn) Repository {
	return &pgRepository{conn: conn}
}
