package spareparts

import (
	"context"
	"database/sql"
	"time"

	"github.com/ferdiebergado/lovemyride/internal/pkg/db"
)

type CreateParams struct {
	Description         string `json:"description"`
	MaintenanceInterval int    `json:"maintenance_interval"`
}

type UpdateParams struct {
	Description         string `json:"description"`
	MaintenanceInterval int    `json:"maintenance_interval"`
}

type sparePartsRepo struct {
	db *sql.DB
}

type Repo interface {
	Create(ctx context.Context, params CreateParams) (*SparePart, error)
	Find(ctx context.Context, id string) (*SparePart, error)
	GetAll(ctx context.Context) ([]SparePart, error)
	Update(ctx context.Context, id string, params UpdateParams) error
	Delete(ctx context.Context, id string) error
}

const (
	maxQueryTime = 60
)

func NewSparePartRepo(db *sql.DB) Repo {
	return &sparePartsRepo{db: db}
}

func (r *sparePartsRepo) Create(ctx context.Context, params CreateParams) (*SparePart, error) {
	const query = `
	INSERT INTO spareparts (description, maintenance_interval)
	VALUES ($1, $2)
	RETURNING id, description, maintenance_interval, created_at, updated_at
	`

	row := r.db.QueryRowContext(ctx, query, params.Description, params.MaintenanceInterval)

	var sparePart SparePart

	err := row.Scan(&sparePart.ID, &sparePart.Description, &sparePart.MaintenanceInterval, &sparePart.CreatedAt, &sparePart.UpdatedAt)

	return &sparePart, err
}

func (r *sparePartsRepo) Find(ctx context.Context, id string) (*SparePart, error) {
	const query = `
	SELECT id, description, maintenance_interval, created_at, updated_at
	FROM spareparts
	WHERE id = $1
	`
	row := r.db.QueryRowContext(ctx, query, id)

	var sparePart SparePart

	err := row.Scan(&sparePart.ID, &sparePart.Description, &sparePart.MaintenanceInterval, &sparePart.CreatedAt, &sparePart.UpdatedAt)

	return &sparePart, err
}

func (r *sparePartsRepo) GetAll(ctx context.Context) ([]SparePart, error) {
	const query = `
	SELECT id, description, maintenance_interval, created_at, updated_at, deleted_at
	FROM spareparts
	`

	ctx, cancel := context.WithTimeout(ctx, maxQueryTime*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, query)

	var spareParts []SparePart

	for rows.Next() {
		var item SparePart

		err = rows.Scan(&item.ID, &item.Description, &item.MaintenanceInterval, &item.CreatedAt, &item.UpdatedAt, &item.DeletedAt)

		if err != nil {
			break
		}

		spareParts = append(spareParts, item)
	}

	// Check for errors during rows "Close".
	// This may be more important if multiple statements are executed
	// in a single batch and rows were written as well as read.
	if closeErr := rows.Close(); closeErr != nil {
		return nil, db.ErrRowClose
	}

	// Check for row scan error.
	if err != nil {
		return nil, db.ErrRowScan
	}

	// Check for errors during row iteration.
	if err = rows.Err(); err != nil {
		return nil, db.ErrRowIteration
	}

	return spareParts, nil
}

func (r *sparePartsRepo) Update(ctx context.Context, id string, params UpdateParams) error {
	const query = `
	UPDATE spareparts
	SET description = $1, maintenance_interval = $2
	WHERE id = $3
	`
	_, err := r.db.ExecContext(ctx, query, params.Description, params.MaintenanceInterval, id)

	return err
}

// Delete implements Repo.
func (r *sparePartsRepo) Delete(ctx context.Context, id string) error {
	const query = `
	UPDATE spareparts
	SET deleted_at = NOW()
	WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, query, id)

	return err
}
