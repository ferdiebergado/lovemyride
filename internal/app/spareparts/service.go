package spareparts

import (
	"context"

	"github.com/ferdiebergado/lovemyride/internal/pkg/db"
)

type service struct {
	repo Repo
}

type Service interface {
	Create(ctx context.Context, params CreateParams) (*SparePart, error)
	Find(ctx context.Context, id string) (*SparePart, error)
	GetAll(ctx context.Context) ([]SparePart, error)
	Update(ctx context.Context, id string, params UpdateParams) error
	Delete(ctx context.Context, id string) error
}

func NewSparePartService(repo Repo) Service {
	return &service{repo: repo}
}

func (s *service) Create(ctx context.Context, params CreateParams) (*SparePart, error) {
	return s.repo.Create(ctx, params)
}

func (s *service) Find(ctx context.Context, id string) (*SparePart, error) {
	return s.repo.Find(ctx, id)
}

func (s *service) GetAll(ctx context.Context) ([]SparePart, error) {
	return s.repo.GetAll(ctx)
}

func (s *service) Update(ctx context.Context, id string, params UpdateParams) error {
	sparePart, err := s.repo.Find(ctx, id)

	if err != nil {
		return err
	}

	if sparePart == nil {
		return db.ErrModelNotFound
	}

	return s.repo.Update(ctx, id, params)
}

func (s *service) Delete(ctx context.Context, id string) error {
	sparePart, err := s.repo.Find(ctx, id)

	if err != nil {
		return err
	}

	if sparePart == nil {
		return db.ErrModelNotFound
	}

	return s.repo.Delete(ctx, id)
}
