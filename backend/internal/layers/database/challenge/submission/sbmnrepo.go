package submission

import "context"

type Repository interface {
	FindAllByInstanceID(ctx context.Context, instId int64) ([]*Submission, error)
}
