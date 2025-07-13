package showcase

import "context"

type Repository interface {
	FindByID(ctx context.Context, id int64) (*ShowcaseChallenge, error)
	FindAll(ctx context.Context) ([]*ShowcaseChallenge, error)
}
