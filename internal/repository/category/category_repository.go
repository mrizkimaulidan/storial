package category

import (
	"context"
	"database/sql"

	"github.com/mrizkimaulidan/storial/internal/entity"
)

type categoryRepository struct {
	//
}

// Find all category on database order by name ASC.
func (cr *categoryRepository) FindAll(ctx context.Context, tx *sql.Tx) (*[]entity.Category, error) {
	query := `
		SELECT
		*
	FROM
		categories
	ORDER BY NAME ASC
	`

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []entity.Category
	for rows.Next() {
		var c entity.Category
		err := rows.Scan(&c.Id, &c.Name, &c.Slug)
		if err != nil {
			return nil, err
		}

		categories = append(categories, c)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &categories, nil
}

func NewRepository() CategoryRepository {
	return &categoryRepository{}
}
