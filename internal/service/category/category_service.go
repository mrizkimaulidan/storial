package category

import (
	"context"
	"database/sql"

	"github.com/mrizkimaulidan/storial/internal/database"
	model "github.com/mrizkimaulidan/storial/internal/model/category"
	"github.com/mrizkimaulidan/storial/internal/repository/category"
	"github.com/mrizkimaulidan/storial/internal/repository/story"
)

type categoryService struct {
	categoryRepository category.CategoryRepository
	storyRepository    story.StoryRepository
	db                 *sql.DB
}

func (cs *categoryService) GetAll(ctx context.Context) (*[]model.CategoryResponse, error) {
	tx, err := cs.db.Begin()
	if err != nil {
		return nil, err
	}
	defer database.CommitOrRollback(tx)

	categories, err := cs.categoryRepository.FindAll(ctx, tx)
	if err != nil {
		return nil, err
	}

	var categoriesResponse []model.CategoryResponse
	for _, c := range *categories {
		storyCounts, err := cs.storyRepository.CountStoryByCategoryID(ctx, tx, c.Id)
		if err != nil {
			return nil, err
		}

		categoryResponse := model.CategoryResponse{
			Id:          c.Id,
			Name:        c.Name,
			Slug:        c.Slug,
			StoryCounts: *storyCounts,
		}

		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	return &categoriesResponse, nil
}

func NewService(cr category.CategoryRepository, sr story.StoryRepository, db *sql.DB) CategoryService {
	return &categoryService{
		categoryRepository: cr,
		storyRepository:    sr,
		db:                 db,
	}
}
