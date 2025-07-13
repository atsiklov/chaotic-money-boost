package showcase

import (
	"context"
	"errors"
	"log"

	myerr "backend/internal/errors"

	"github.com/jackc/pgx/v5"
)

const query = `
	select ci.id, ct.category, ct.description, ci.started_at, ci.expires_at
	from challenge_instances ci join challenge_templates ct on ci.challenge_template_id = ct.id
`

type pgRepository struct {
	conn *pgx.Conn
}

func (repo *pgRepository) FindByID(ctx context.Context, id int64) (*ShowcaseChallenge, error) {
	var chgeShow ShowcaseChallenge
	err := repo.conn.QueryRow(ctx, query+" where ci.id = $1", id).
		Scan(&chgeShow.InstanceID, &chgeShow.Category, &chgeShow.Description, &chgeShow.StartedAt, &chgeShow.ExpiresAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Printf("No showcase challenge with id=%d", id)
			return nil, myerr.ErrChgeShowcaseNotFound
		}
		log.Printf("Failed to find showcase challenge with id=%d: %s", id, err.Error())
		return nil, err
	}
	return &chgeShow, nil
}

func (repo *pgRepository) FindAll(ctx context.Context) ([]*ShowcaseChallenge, error) {
	rows, err := repo.conn.Query(ctx, query)
	if err != nil {
		log.Printf("Failed to select challenges: %s", err.Error())
		return nil, err
	}
	defer rows.Close()

	var chgesShow []*ShowcaseChallenge
	for rows.Next() {
		var chgeShow ShowcaseChallenge
		if err := rows.Scan(
			&chgeShow.InstanceID,
			&chgeShow.Category,
			&chgeShow.Description,
			&chgeShow.StartedAt,
			&chgeShow.ExpiresAt,
		); err != nil {
			log.Printf("Failed to parse row into challenge showcase: %s", err.Error())
			return nil, err
		}
		chgesShow = append(chgesShow, &chgeShow)
	}
	rows.Close() // дважды намеренно
	if err := rows.Err(); err != nil {
		log.Printf("Error during iteration over rows: %s", err.Error())
		return nil, err
	}
	return chgesShow, nil
}

func NewPgRepo(conn *pgx.Conn) Repository {
	return &pgRepository{conn: conn}
}
