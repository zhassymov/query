package query

import (
	"context"

	"github.com/zhassymov/please"
)

type Repository[T any] interface {
	Create(ctx context.Context, entity T) error
	Update(ctx context.Context, entity T) error
	Delete(ctx context.Context, entity T) error

	Find(ctx context.Context, filters ...please.Validate[*Query]) ([]T, error)
	Count(ctx context.Context, filters ...please.Validate[*Query]) (int64, error)
}

type EventStore[T any] interface {
	Append(ctx context.Context, version int64, event T) error

	Fetch(ctx context.Context, filters ...please.Validate[*Query]) ([]T, error)
	Count(ctx context.Context, filters ...please.Validate[*Query]) (int64, error)
}
