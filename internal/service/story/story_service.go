package story

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/mrizkimaulidan/storial/internal/database"
	"github.com/mrizkimaulidan/storial/internal/entity"
	categorymodel "github.com/mrizkimaulidan/storial/internal/model/category"
	model "github.com/mrizkimaulidan/storial/internal/model/story"
	usermodel "github.com/mrizkimaulidan/storial/internal/model/user"
	"github.com/mrizkimaulidan/storial/internal/repository/chapter"
	"github.com/mrizkimaulidan/storial/internal/repository/story"
	chapterservice "github.com/mrizkimaulidan/storial/internal/service/chapter"
	"github.com/mrizkimaulidan/storial/internal/service/file"
	exception "github.com/mrizkimaulidan/storial/pkg/exception/story"
	"github.com/mrizkimaulidan/storial/pkg/time"
)

var (
	COVER_PATH = "public/cover"
)

type storyService struct {
	storyRepository   story.StoryRepository
	fileService       file.FileService
	chapterRepository chapter.ChapterRepository
	chapterService    chapterservice.ChapterService
	db                *sql.DB
}

func (ss *storyService) ProcessStoryResponseFilter(ctx context.Context, tx *sql.Tx, stories []entity.Story) (*[]model.StoryResponseByFilter, error) {
	var storiesResponse []model.StoryResponseByFilter
	for _, s := range stories {
		chapterCounts, err := ss.chapterRepository.CountChapterByStorySlug(ctx, tx, s.Slug)
		if err != nil {
			return nil, err
		}

		sr := model.StoryResponseByFilter{
			Id:     s.Id,
			UserID: s.UserID,
			User: usermodel.UserResponseByFilter{
				Id:       s.User.Id,
				Name:     s.User.Name,
				Username: s.User.Username,
			},
			Title:         s.Title,
			Slug:          s.Slug,
			ChapterCounts: *chapterCounts,
			IsAdult:       s.IsAdult,
			Cover:         s.CoverPath(),
			CreatedAt:     time.UnixToTime(s.CreatedAt),
		}

		storiesResponse = append(storiesResponse, sr)
	}

	return &storiesResponse, nil
}

func (ss *storyService) ProcessStoryResponseFilterByCategorySlug(ctx context.Context, tx *sql.Tx, stories []entity.Story) (*[]model.StoryResponseByCategorySlug, error) {
	var storiesResponse []model.StoryResponseByCategorySlug
	for _, s := range stories {
		chapterCounts, err := ss.chapterRepository.CountChapterByStorySlug(ctx, tx, s.Slug)
		if err != nil {
			return nil, err
		}

		storyResponse := model.StoryResponseByCategorySlug{
			Id:     s.Id,
			UserID: s.UserID,
			User: usermodel.UserResponseBySlug{
				Id:       s.User.Id,
				Name:     s.User.Name,
				Username: s.User.Username,
			},
			Title:         s.Title,
			Slug:          s.Slug,
			ChapterCounts: *chapterCounts,
			CreatedAt:     time.UnixToTime(s.CreatedAt),
			UpdatedAt:     time.UnixToTime(s.UpdatedAt),
		}

		storiesResponse = append(storiesResponse, storyResponse)
	}

	return &storiesResponse, nil
}

func (ss *storyService) FilterStoryByCategorySlug(ctx context.Context, categorySlug string, filterType string) (*[]model.StoryResponseByCategorySlug, error) {
	tx, err := ss.db.Begin()
	if err != nil {
		return nil, err
	}
	defer database.CommitOrRollback(tx)

	switch filterType {
	case "time":
		stories, err := ss.storyRepository.FilterLatestBasedOnCategorySlug(ctx, tx, categorySlug)
		if err != nil {
			return nil, err
		}

		storiesResponse, err := ss.ProcessStoryResponseFilterByCategorySlug(ctx, tx, *stories)
		if err != nil {
			return nil, err
		}

		return storiesResponse, nil
	case "modified":
		stories, err := ss.storyRepository.FilterLatestModifiedChapterByCategorySlug(ctx, tx, categorySlug)
		if err != nil {
			return nil, err
		}

		storiesResponse, err := ss.ProcessStoryResponseFilterByCategorySlug(ctx, tx, *stories)
		if err != nil {
			return nil, err
		}

		return storiesResponse, nil
	default:
		stories, err := ss.storyRepository.FindByCategorySlug(ctx, tx, categorySlug)
		if err != nil {
			return nil, err
		}

		storiesResponse, err := ss.ProcessStoryResponseFilterByCategorySlug(ctx, tx, *stories)
		if err != nil {
			return nil, err
		}

		return storiesResponse, nil
	}
}

func (ss *storyService) GetStoryByCategorySlug(ctx context.Context, categorySlug string) (*[]model.StoryResponseByCategorySlug, error) {
	tx, err := ss.db.Begin()
	if err != nil {
		return nil, err
	}
	defer database.CommitOrRollback(tx)

	stories, err := ss.storyRepository.FindByCategorySlug(ctx, tx, categorySlug)
	if err != nil {
		return nil, err
	}

	storiesResponse, err := ss.ProcessStoryResponseFilterByCategorySlug(ctx, tx, *stories)
	if err != nil {
		return nil, err
	}

	return storiesResponse, nil
}

func (ss *storyService) GetAllStory(ctx context.Context, userID uint64) (*[]model.StoryResponse, error) {
	tx, err := ss.db.Begin()
	if err != nil {
		return nil, err
	}
	defer database.CommitOrRollback(tx)

	stories, err := ss.storyRepository.FindAllByUserID(ctx, tx, userID)
	if err != nil {
		return nil, err
	}

	var storiesResponse []model.StoryResponse
	for _, s := range *stories {
		storyResponse := model.StoryResponse{
			Id:     s.Id,
			UserID: s.UserID,
			User: usermodel.UserResponse{
				Id:    s.User.Id,
				Name:  s.User.Name,
				Email: s.User.Email,
			},
			Title:       s.Title,
			Slug:        s.Slug,
			IsPublished: s.IsPublished,
			CreatedAt:   time.UnixToTime(s.CreatedAt),
			UpdatedAt:   time.UnixToTime(s.UpdatedAt),
		}

		storiesResponse = append(storiesResponse, storyResponse)
	}

	return &storiesResponse, nil
}

func (ss *storyService) RemoveStory(ctx context.Context, r model.DeleteStoryRequest) (*model.DeletetedStoryResponse, error) {
	tx, err := ss.db.Begin()
	if err != nil {
		return nil, err
	}
	defer database.CommitOrRollback(tx)

	id, err := strconv.Atoi(r.Id)
	if err != nil {
		return nil, err
	}

	story, err := ss.storyRepository.FindByID(ctx, tx, uint64(id))
	if err != nil {
		return nil, err
	}

	err = ss.storyRepository.Delete(ctx, tx, uint64(id))
	if err != nil {
		return nil, err
	}

	ss.fileService.RemoveFile(story.Cover)

	return &model.DeletetedStoryResponse{
		Status: true,
	}, nil
}

func (ss *storyService) LoadStoryImageCover(ctx context.Context, filename string) ([]byte, error) {
	ss.fileService.SetFilename(filename)
	fullPath := ss.fileService.GetFullPath()

	bytes, err := ss.fileService.Get(fullPath)
	if err != nil {
		return nil, exception.ErrCoverImageNotFound
	}

	return bytes, nil
}

func (ss *storyService) EditStory(ctx context.Context, r model.UpdateStoryRequest) (*model.UpdatedStoryResponse, error) {
	tx, err := ss.db.Begin()
	if err != nil {
		return nil, err
	}
	defer database.CommitOrRollback(tx)

	categoryID, err := strconv.Atoi(r.CategoryID)
	if err != nil {
		return nil, err
	}

	isAdult, err := strconv.ParseBool(r.IsAdult)
	if err != nil {
		return nil, err
	}

	isPublished, err := strconv.ParseBool(r.IsPublished)
	if err != nil {
		return nil, err
	}

	story, err := ss.storyRepository.FindBySlugAndUserID(ctx, tx, r.Slug, r.UserID)
	if err != nil {
		return nil, err
	}

	var s entity.Story
	s = entity.Story{
		Id:          story.Id,
		UserID:      r.UserID,
		CategoryID:  uint64(categoryID),
		Title:       r.Title,
		Slug:        fmt.Sprintf("%s-%d", s.ToSlug(r.Title), story.Id),
		Description: r.Description,
		IsAdult:     isAdult,
		IsPublished: isPublished,
		UpdatedAt:   time.CurrentTimeToUnixTimestamp(),
	}

	// upload file if file exists on request struct
	if r.Cover != nil {
		filename, err := ss.fileService.Upload(r.Cover, r.CoverFileheader)
		if err != nil {
			return nil, err
		}

		s.Cover = filename

		// remove old cover file
		ss.fileService.RemoveFile(story.Cover)
	}

	updatedStory, err := ss.storyRepository.Update(ctx, tx, r.Slug, s)
	if err != nil {
		return nil, err
	}

	return &model.UpdatedStoryResponse{
		Id:         updatedStory.Id,
		UserID:     updatedStory.UserID,
		CategoryID: updatedStory.CategoryID,
		Category: categorymodel.CategoryResponse{
			Id:   updatedStory.Category.Id,
			Slug: updatedStory.Category.Slug,
			Name: updatedStory.Category.Name,
		},
		Title:       updatedStory.Title,
		Slug:        updatedStory.Slug,
		Description: updatedStory.Description,
		IsAdult:     updatedStory.IsAdult,
		IsPublished: updatedStory.IsPublished,
		Cover:       updatedStory.CoverPath(),
		CreatedAt:   time.UnixToTime(updatedStory.CreatedAt),
		UpdatedAt:   time.UnixToTime(updatedStory.UpdatedAt),
	}, nil
}

func (ss *storyService) AddStory(ctx context.Context, r model.CreateStoryRequest) (*model.CreatedStoryResponse, error) {
	tx, err := ss.db.Begin()
	if err != nil {
		return nil, err
	}
	defer database.CommitOrRollback(tx)

	err = r.Validate()
	if err != nil {
		return nil, err
	}

	var story entity.Story
	categoryID, err := strconv.Atoi(r.CategoryID)
	if err != nil {
		return nil, err
	}

	isAdult, err := strconv.ParseBool(r.IsAdult)
	if err != nil {
		return nil, err
	}

	isPublished, err := strconv.ParseBool(r.IsPublished)
	if err != nil {
		return nil, err
	}

	id := uint64(story.GenerateID())
	story = entity.Story{
		Id:          id,
		UserID:      r.UserID,
		CategoryID:  uint64(categoryID),
		Title:       r.Title,
		Slug:        fmt.Sprintf("%s-%d", story.ToSlug(r.Title), id),
		Description: r.Description,
		IsAdult:     isAdult,
		IsPublished: isPublished,
		CreatedAt:   time.CurrentTimeToUnixTimestamp(),
		UpdatedAt:   time.CurrentTimeToUnixTimestamp(),
	}

	// upload file if file exists on request struct
	if r.Cover != nil {
		filename, err := ss.fileService.Upload(r.Cover, r.CoverFileheader) // upload file to local system
		if err != nil {
			return nil, err
		}

		story.Cover = filename
	}

	createdStory, err := ss.storyRepository.Save(ctx, tx, story)
	if err != nil {
		return nil, err
	}

	return &model.CreatedStoryResponse{
		Id:         createdStory.Id,
		UserID:     createdStory.UserID,
		CategoryID: createdStory.CategoryID,
		Category: categorymodel.CategoryResponse{
			Id:   createdStory.Category.Id,
			Slug: createdStory.Category.Slug,
			Name: createdStory.Category.Name,
		},
		Title:       createdStory.Title,
		Slug:        createdStory.Slug,
		Description: createdStory.Description,
		IsAdult:     createdStory.IsAdult,
		IsPublished: createdStory.IsPublished,
		Cover:       createdStory.CoverPath(),
		CreatedAt:   time.UnixToTime(createdStory.CreatedAt),
		UpdatedAt:   time.UnixToTime(createdStory.UpdatedAt),
	}, nil
}

func (ss *storyService) GetStoryBySlug(ctx context.Context, slug string) (*model.StoryResponseBySlug, error) {
	tx, err := ss.db.Begin()
	if err != nil {
		return nil, err
	}
	defer database.CommitOrRollback(tx)

	story, err := ss.storyRepository.FindBySlug(ctx, tx, slug)
	if err != nil {
		return nil, err
	}

	counts, err := ss.chapterRepository.CountChapterByStorySlug(ctx, tx, story.Slug)
	if err != nil {
		return nil, err
	}

	chapters, err := ss.chapterRepository.FindAllChapterByStorySlug(ctx, tx, story.Slug)
	if err != nil {
		return nil, err
	}

	readingTime, err := ss.chapterService.CalculateReadingTimeByChapters(ctx, *chapters)
	if err != nil {
		return nil, err
	}

	return &model.StoryResponseBySlug{
		Id:     story.Id,
		UserID: story.UserID,
		User: usermodel.UserResponseBySlug{
			Id:       story.User.Id,
			Name:     story.User.Name,
			Username: story.User.Username,
		},
		CategoryID: story.CategoryID,
		Category: categorymodel.CategoryResponseByStorySlug{
			Id:   story.Category.Id,
			Name: story.Category.Name,
			Slug: story.Category.Slug,
		},
		Title:         story.Title,
		Slug:          story.Slug,
		ChapterCounts: *counts,
		ReadingTime:   readingTime,
		CreatedAt:     time.UnixToTime(story.CreatedAt),
		UpdatedAt:     time.UnixToTime(story.UpdatedAt),
	}, nil
}

func (ss *storyService) FilterStory(ctx context.Context, filterType string) (*[]model.StoryResponseByFilter, error) {
	tx, err := ss.db.Begin()
	if err != nil {
		return nil, err
	}
	defer database.CommitOrRollback(tx)

	switch filterType {
	case "time":
		stories, err := ss.storyRepository.FilterLatest(ctx, tx)
		if err != nil {
			return nil, err
		}

		storiesResponse, err := ss.ProcessStoryResponseFilter(ctx, tx, *stories)
		if err != nil {
			return nil, err
		}

		return storiesResponse, nil
	case "modified":
		stories, err := ss.storyRepository.FilterLatestModifiedChapter(ctx, tx)
		if err != nil {
			return nil, err
		}

		storiesResponse, err := ss.ProcessStoryResponseFilter(ctx, tx, *stories)
		if err != nil {
			return nil, err
		}

		return storiesResponse, nil
	default:
		stories, err := ss.storyRepository.FilterLatestModifiedChapter(ctx, tx)
		if err != nil {
			return nil, err
		}

		storiesResponse, err := ss.ProcessStoryResponseFilter(ctx, tx, *stories)
		if err != nil {
			return nil, err
		}

		return storiesResponse, nil
	}
}

func NewService(storyRepository story.StoryRepository, cr chapter.ChapterRepository, cs chapterservice.ChapterService, db *sql.DB) StoryService {
	return &storyService{
		storyRepository:   storyRepository,
		fileService:       file.NewService(COVER_PATH),
		chapterRepository: cr,
		chapterService:    cs,
		db:                db,
	}
}
