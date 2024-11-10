package servicelogs

import (
	"context"
	"database/sql"

	"github.com/ferdiebergado/lovemyride/internal/pkg/db"
)

type Repo struct {
	db *sql.DB
}

func NewServiceLogsRepo(db *sql.DB) db.Repository[*ServiceLog] {
	return &Repo{db: db}
}

// Create implements db.Repository.
func (r *Repo) Create(context.Context, any) (*ServiceLog, error) {
	panic("unimplemented")
}

// Delete implements db.Repository.
func (r *Repo) Delete(context.Context, string) error {
	panic("unimplemented")
}

// Find implements db.Repository.
func (r *Repo) Find(context.Context, string) (*ServiceLog, error) {
	panic("unimplemented")
}

// FindAll implements db.Repository.
func (r *Repo) FindAll(context.Context, any) ([]*ServiceLog, error) {
	panic("unimplemented")
}

// Update implements db.Repository.
func (r *Repo) Update(context.Context, any) error {
	panic("unimplemented")
}
