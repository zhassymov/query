package query

import (
	"context"
	"github.com/zhassymov/please"
)

type Repository[T any] interface {
	Create(context.Context, T) error
	Update(context.Context, T) error
	Delete(context.Context, T) error

	Find(context.Context, ...please.Validate[*Query])
	Count(context.Context, ...please.Validate[*Query])
}
