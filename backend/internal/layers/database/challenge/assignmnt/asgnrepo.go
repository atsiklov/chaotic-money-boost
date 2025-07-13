package assignment

import "context"

type Repository interface {
	Create(ctx context.Context, asgnNew *Assignment) error
	Update(ctx context.Context, asgnUpd *Assignment) error
}
