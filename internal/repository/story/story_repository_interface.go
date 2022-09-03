package story

import (
	"context"
	"database/sql"

	"github.com/mrizkimaulidan/storial/internal/entity"
)

type StoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, s entity.Story) (*entity.Story, error)
	Update(ctx context.Context, tx *sql.Tx, slug string, s entity.Story) (*entity.Story, error)
	FindBySlugAndUserID(ctx context.Context, tx *sql.Tx, slug string, userID uint64) (*entity.Story, error)
	FindByID(ctx context.Context, tx *sql.Tx, id uint64) (*entity.Story, error)
	Delete(ctx context.Context, tx *sql.Tx, id uint64) error
	FindAllByUserID(ctx context.Context, tx *sql.Tx, userID uint64) (*[]entity.Story, error)
	FindBySlug(ctx context.Context, tx *sql.Tx, slug string) (*entity.Story, error)
	FilterLatest(ctx context.Context, tx *sql.Tx) (*[]entity.Story, error)
	FilterLatestModifiedChapter(ctx context.Context, tx *sql.Tx) (*[]entity.Story, error)
	CountStoryByCategoryID(ctx context.Context, tx *sql.Tx, categoryID uint64) (*uint64, error)
	FindByCategorySlug(ctx context.Context, tx *sql.Tx, categorySlug string) (*[]entity.Story, error)
	FilterLatestBasedOnCategorySlug(ctx context.Context, tx *sql.Tx, categorySlug string) (*[]entity.Story, error)
	FilterLatestModifiedChapterByCategorySlug(ctx context.Context, tx *sql.Tx, categorySlug string) (*[]entity.Story, error)
}
