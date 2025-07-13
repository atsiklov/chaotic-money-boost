package instance

import "context"

type Repository interface {
	Create(ctx context.Context, chgeInst *ChallengeInstance) (*ChallengeInstance, error)
}
