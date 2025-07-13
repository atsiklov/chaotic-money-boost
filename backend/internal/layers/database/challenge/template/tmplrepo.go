package template

import "context"

type Repository interface {
	FindByID(ctx context.Context, id int64) (*ChallengeTemplate, error)
}
