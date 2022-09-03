package chapter

import (
	"context"

	"github.com/mrizkimaulidan/storial/internal/entity"
	model "github.com/mrizkimaulidan/storial/internal/model/chapter"
)

type ChapterService interface {
	AddChapter(ctx context.Context, r model.CreateChapterRequest) (*model.CreatedChapterdResponse, error)
	EditChapter(ctx context.Context, r model.UpdateChapterRequest) (*model.UpdatedChapterResponse, error)
	GetChapterByStorySlugAndChapterSlug(ctx context.Context, userID uint64, storySlug string, chapterSlug string) (*model.ChapterResponseByStorySlugAndChapterSlug, error)
	RemoveChapter(ctx context.Context, userID uint64, chapterID string) (*model.DeletedChapterResponse, error)
	CalculateReadingTimeByChapters(ctx context.Context, chapters []entity.Chapter) (string, error)
	GetAllChapterByStorySlug(ctx context.Context, storySlug string) (*[]model.ChapterResponseBySlug, error)
	GetAllChapterByStoryID(ctx context.Context, storyID string) (*[]model.ChapterResponseBySlug, error)
	LikeChapter(ctx context.Context, storyID string, chapterID string, userID uint64) (*model.LikedChapterResponse, error)
}
