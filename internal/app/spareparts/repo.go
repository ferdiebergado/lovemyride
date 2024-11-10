package spareparts

import (
	"context"
	"database/sql"

	"github.com/ferdiebergado/lovemyride/internal/pkg/db"
)

type Repo struct {
	db *sql.DB
}

// Create implements db.Repository.
func (s *Repo) Create(ctx context.Context, params any) (*SparePart, error) {
	panic("unimplemented")
}

// Delete implements db.Repository.
func (s *Repo) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// Find implements db.Repository.
func (s *Repo) Find(ctx context.Context, id string) (*SparePart, error) {
	panic("unimplemented")
}

// FindAll implements db.Repository.
func (s *Repo) FindAll(ctx context.Context, filter any) ([]*SparePart, error) {
	panic("unimplemented")
}

// Update implements db.Repository.
func (s *Repo) Update(ctx context.Context, params any) error {
	panic("unimplemented")
}

func NewSparePartsRepo(db *sql.DB) db.Repository[*SparePart] {
	return &Repo{db: db}
}
