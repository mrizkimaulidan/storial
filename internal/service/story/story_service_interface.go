package story

import (
	"context"
	"database/sql"

	"github.com/mrizkimaulidan/storial/internal/entity"
	model "github.com/mrizkimaulidan/storial/internal/model/story"
)

type StoryService interface {
	AddStory(ctx context.Context, r model.CreateStoryRequest) (*model.CreatedStoryResponse, error)
	EditStory(ctx context.Context, r model.UpdateStoryRequest) (*model.UpdatedStoryResponse, error)
	LoadStoryImageCover(ctx context.Context, filename string) ([]byte, error)
	RemoveStory(ctx context.Context, r model.DeleteStoryRequest) (*model.DeletetedStoryResponse, error)
	GetAllStory(ctx context.Context, userID uint64) (*[]model.StoryResponse, error)
	GetStoryBySlug(ctx context.Context, slug string) (*model.StoryResponseBySlug, error)
	FilterStory(ctx context.Context, filterType string) (*[]model.StoryResponseByFilter, error)
	GetStoryByCategorySlug(ctx context.Context, categorySlug string) (*[]model.StoryResponseByCategorySlug, error)
	FilterStoryByCategorySlug(ctx context.Context, categorySlug string, filterType string) (*[]model.StoryResponseByCategorySlug, error)
	ProcessStoryResponseFilterByCategorySlug(ctx context.Context, tx *sql.Tx, stories []entity.Story) (*[]model.StoryResponseByCategorySlug, error)
	ProcessStoryResponseFilter(ctx context.Context, tx *sql.Tx, stories []entity.Story) (*[]model.StoryResponseByFilter, error)
}
