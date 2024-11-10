package spareparts

import "github.com/ferdiebergado/lovemyride/internal/pkg/db"

type SparePart struct {
	db.Model
	Description         string `json:"description"`
	MaintenanceInterval int    `json:"maintenance_interval"`
}
