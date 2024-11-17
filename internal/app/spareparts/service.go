package spareparts

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ferdiebergado/lovemyride/internal/pkg/db"
)

type service struct {
	repo Repo
}

type Service interface {
	Create(ctx context.Context, params CreateParams) (SparePart, error)
	Find(ctx context.Context, id string) (SparePart, error)
	GetAll(ctx context.Context) ([]SparePart, error)
	Update(ctx context.Context, id string, params UpdateParams) error
	Delete(ctx context.Context, id string, mode db.DeleteMode) error
}

func NewSparePartService(repo Repo) Service {
	return &service{repo: repo}
}

func (s *service) Create(ctx context.Context, params CreateParams) (SparePart, error) {
	return s.repo.Create(ctx, params)
}

func (s *service) Find(ctx context.Context, id string) (SparePart, error) {
	sparePart, err := s.repo.Find(ctx, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return SparePart{}, db.ErrModelNotFound
		}

		return SparePart{}, err
	}

	return sparePart, nil
}

func (s *service) GetAll(ctx context.Context) ([]SparePart, error) {
	return s.repo.GetAll(ctx)
}

func (s *service) Update(ctx context.Context, id string, params UpdateParams) error {
	_, err := s.repo.Find(ctx, id)

	if err != nil {
		return err
	}

	return s.repo.Update(ctx, id, params)
}

func (s *service) Delete(ctx context.Context, id string, mode db.DeleteMode) error {
	_, err := s.repo.Find(ctx, id)

	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, id, mode)
}
