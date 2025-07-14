package submission

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type pgRepository struct {
	conn *pgx.Conn
}

func (repo *pgRepository) FindAllByInstanceID(ctx context.Context, instID int64) ([]*Submission, error) {
	const query = `select user_id, submission, submitted_at from challenge_assignments where challenge_instance_id = $1;`
	rows, err := repo.conn.Query(ctx, query, instID)
	if err != nil {
		log.Printf("Failed to select submissions: %s", err.Error())
		return nil, err
	}
	defer rows.Close()

	var submissions []*Submission
	for rows.Next() {
		var submission Submission
		if err := rows.Scan(&submission.UserID, &submission.Submission, &submission.SubmittedAt); err != nil {
			log.Printf("Failed to parse row into submission: %s", err.Error())
			return nil, err
		}
		submissions = append(submissions, &submission)
	}
	rows.Close() // дважды намеренно
	if err := rows.Err(); err != nil {
		log.Printf("Error during iteration over rows: %s", err.Error())
		return nil, err
	}
	return submissions, nil
}

func NewPgRepo(conn *pgx.Conn) Repository {
	return &pgRepository{conn: conn}
}
