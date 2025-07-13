package instance

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type pgRepository struct {
	conn *pgx.Conn
}

func (repo *pgRepository) Create(ctx context.Context, chgeInstNew *ChallengeInstance) (*ChallengeInstance, error) {
	const query = `
	    insert into challenge_instances (challenge_template_id, status) values ($1, $2)
	    returning id, challenge_template_id, status, created_at, updated_at, started_at, expires_at
	`
	var chgeInstSaved ChallengeInstance
	err := repo.conn.QueryRow(ctx, query,
		chgeInstNew.TemplateID,
		chgeInstNew.Status,
	).Scan(
		&chgeInstSaved.ID,
		&chgeInstSaved.TemplateID,
		&chgeInstSaved.Status,
		&chgeInstSaved.CreatedAt,
		&chgeInstSaved.UpdatedAt,
		&chgeInstSaved.StartedAt,
		&chgeInstSaved.ExpiresAt,
	)

	if err != nil {
		log.Printf("Failed to create new instance with status=%s: %s", chgeInstNew.Status, err.Error())
		return nil, err
	}
	log.Printf("Created new instance with id = %d and status = %s", chgeInstSaved.ID, chgeInstSaved.Status)
	return &chgeInstSaved, nil
}

func NewPgRepo(conn *pgx.Conn) Repository {
	return &pgRepository{conn: conn}
}
