package chapter

import (
	"context"
	"database/sql"

	"github.com/mrizkimaulidan/storial/internal/entity"
)

type ChapterRepository interface {
	Save(ctx context.Context, tx *sql.Tx, c entity.Chapter) (*entity.Chapter, error)
	Update(ctx context.Context, tx *sql.Tx, userID uint64, storySlug string, chapterSlug string, c entity.Chapter) (*entity.Chapter, error)
	FindByStorySlugAndChapterSlug(ctx context.Context, tx *sql.Tx, userID uint64, storySlug string, chapterSlug string) (*entity.Chapter, error)
	Delete(ctx context.Context, tx *sql.Tx, userID uint64, chapterID uint64) error
	FindByID(ctx context.Context, tx *sql.Tx, id uint64) (*entity.Chapter, error)
	CountChapterByStorySlug(ctx context.Context, tx *sql.Tx, storySlug string) (*uint64, error)
	FindAllChapterByStorySlug(ctx context.Context, tx *sql.Tx, storySlug string) (*[]entity.Chapter, error)
	FindAllChapterByStoryID(ctx context.Context, tx *sql.Tx, storyID uint64) (*[]entity.Chapter, error)
	CountChapterLikesByChapterID(ctx context.Context, tx *sql.Tx, chapterID uint64) (*uint64, error)
	CountChapterByUserID(ctx context.Context, tx *sql.Tx, userID uint64) (*uint64, error)
	SaveChapterLikes(ctx context.Context, tx *sql.Tx, chapterID uint64, userID uint64) error
}
