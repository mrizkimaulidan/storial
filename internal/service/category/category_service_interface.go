package category

import (
	"context"

	model "github.com/mrizkimaulidan/storial/internal/model/category"
)

type CategoryService interface {
	GetAll(ctx context.Context) (*[]model.CategoryResponse, error)
}
