package servicelogs

import (
	"database/sql"
)

type Repo struct {
	db *sql.DB
}
