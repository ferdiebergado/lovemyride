package db

import "context"

type Repository[T any] interface {
	Create(ctx context.Context, params any) (T, error)
	Find(ctx context.Context, id string) (T, error)
	Update(ctx context.Context, params any) error
	Delete(ctx context.Context, id string) error
	FindAll(ctx context.Context, filter any) ([]T, error)
}

type Finder[T any] interface {
	Find(ctx context.Context, id string) (*T, error)
}
