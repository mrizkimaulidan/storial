package category

import (
	"context"
	"database/sql"

	"github.com/mrizkimaulidan/storial/internal/entity"
)

type CategoryRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) (*[]entity.Category, error)
}
