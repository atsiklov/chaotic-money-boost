package template

import (
	myerr "backend/internal/errors"
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
)

type pgRepository struct {
	conn *pgx.Conn
}

func (r *pgRepository) FindByID(ctx context.Context, id int64) (*ChallengeTemplate, error) {
	query := `select id, category, description, duration from challenge_templates where id = $1`
	row := r.conn.QueryRow(ctx, query, id)

	var chgeTmpl ChallengeTemplate
	err := row.Scan(&chgeTmpl.ID, &chgeTmpl.Category, &chgeTmpl.Description, &chgeTmpl.Duration)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Printf("No challenge template with id=%d", id)
			return nil, myerr.ErrChgeTemplateNotFound
		}
		log.Printf("Failed to find challenge template with id=%d: %s", id, err.Error())
		return nil, err
	}
	return &chgeTmpl, nil
}

func NewPgRepo(conn *pgx.Conn) Repository {
	return &pgRepository{conn: conn}
}
