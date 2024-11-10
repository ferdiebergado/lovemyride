package db

import "context"

type Repository[T any] interface {
	Create(context.Context, any) (T, error)
	Find(context.Context, string) (T, error)
	Update(context.Context, any) error
	Delete(context.Context, string) error
	FindAll(context.Context, any) ([]T, error)
}

type Finder[T any] interface {
	Find(ctx context.Context, id string) (*T, error)
}
