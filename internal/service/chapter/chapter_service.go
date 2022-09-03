package chapter

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/mrizkimaulidan/storial/internal/database"
	"github.com/mrizkimaulidan/storial/internal/entity"
	model "github.com/mrizkimaulidan/storial/internal/model/chapter"
	storymodel "github.com/mrizkimaulidan/storial/internal/model/story"
	usermodel "github.com/mrizkimaulidan/storial/internal/model/user"
	"github.com/mrizkimaulidan/storial/internal/repository/chapter"
	"github.com/mrizkimaulidan/storial/internal/repository/story"
	exception "github.com/mrizkimaulidan/storial/pkg/exception/chapter"
	"github.com/mrizkimaulidan/storial/pkg/time"
)

type chapterService struct {
	chapterRepository chapter.ChapterRepository
	storyRepository   story.StoryRepository
	db                *sql.DB
}

func (cs *chapterService) LikeChapter(ctx context.Context, storyID string, chapterID string, userID uint64) (*model.LikedChapterResponse, error) {
	tx, err := cs.db.Begin()
	if err != nil {
		return nil, err
	}
	defer database.CommitOrRollback(tx)

	sId, err := strconv.Atoi(storyID)
	if err != nil {
		return nil, err
	}

	_, err = cs.storyRepository.FindByID(ctx, tx, uint64(sId))
	if err != nil {
		return nil, err
	}

	cId, err := strconv.Atoi(chapterID)
	if err != nil {
		return nil, err
	}

	chapter, err := cs.chapterRepository.FindByID(ctx, tx, uint64(cId))
	if err != nil {
		return nil, err
	}

	if chapter.Story.UserID == userID {
		return nil, exception.ErrCannotLikeYourOwnChapter
	}

	err = cs.chapterRepository.SaveChapterLikes(ctx, tx, uint64(cId), userID)
	if err != nil {
		return nil, err
	}

	return &model.LikedChapterResponse{
		Status: true,
	}, nil
}

func (cs *chapterService) AddChapter(ctx context.Context, r model.CreateChapterRequest) (*model.CreatedChapterdResponse, error) {
	tx, err := cs.db.Begin()
	if err != nil {
		return nil, err
	}
	defer database.CommitOrRollback(tx)

	err = r.Validate()
	if err != nil {
		return nil, err
	}

	isPublished, err := strconv.ParseBool(r.IsPublished)
	if err != nil {
		return nil, err
	}

	story, err := cs.storyRepository.FindBySlugAndUserID(ctx, tx, r.StorySlug, r.UserID)
	if err != nil {
		return nil, err
	}

	var c entity.Chapter
	c = entity.Chapter{
		Id:            uint64(c.GenerateID()),
		StoryID:       story.Id,
		Title:         r.Title,
		Slug:          c.ToSlug(r.Title),
		Body:          r.Body,
		AuthorComment: r.AuthorComment,
		IsPublished:   isPublished,
		CreatedAt:     time.CurrentTimeToUnixTimestamp(),
		UpdatedAt:     time.CurrentTimeToUnixTimestamp(),
	}

	c.WordCounts = uint64(c.CountChars())
	c.ReadingTime = c.CalculateReadingTime()

	createdChapter, err := cs.chapterRepository.Save(ctx, tx, c)
	if err != nil {
		return nil, err
	}

	return &model.CreatedChapterdResponse{
		Id:            createdChapter.Id,
		StoryID:       createdChapter.StoryID,
		Title:         createdChapter.Title,
		Slug:          createdChapter.Slug,
		Body:          createdChapter.Body,
		AuthorComment: createdChapter.AuthorComment,
		WordCounts:    createdChapter.WordCounts,
		ReadingTime:   createdChapter.ReadingTime,
		IsPublished:   createdChapter.IsPublished,
		CreatedAt:     time.UnixToTime(createdChapter.CreatedAt),
		UpdatedAt:     time.UnixToTime(createdChapter.UpdatedAt),
	}, nil
}

func (cs *chapterService) EditChapter(ctx context.Context, r model.UpdateChapterRequest) (*model.UpdatedChapterResponse, error) {
	tx, err := cs.db.Begin()
	if err != nil {
		return nil, err
	}
	defer database.CommitOrRollback(tx)

	err = r.Validate()
	if err != nil {
		return nil, err
	}

	story, err := cs.storyRepository.FindBySlug(ctx, tx, r.StorySlug)
	if err != nil {
		return nil, err
	}

	_, err = cs.chapterRepository.FindByStorySlugAndChapterSlug(ctx, tx, r.UserID, r.StorySlug, r.ChapterSlug)
	if err != nil {
		return nil, err
	}

	isPublished, err := strconv.ParseBool(r.IsPublished)
	if err != nil {
		return nil, err
	}

	var c entity.Chapter
	c = entity.Chapter{
		StoryID:       story.Id,
		Title:         r.Title,
		Slug:          c.ToSlug(r.Title),
		Body:          r.Body,
		AuthorComment: r.AuthorComment,
		IsPublished:   isPublished,
		UpdatedAt:     time.CurrentTimeToUnixTimestamp(),
	}

	c.WordCounts = uint64(c.CountChars())
	c.ReadingTime = c.CalculateReadingTime()

	updatedChapter, err := cs.chapterRepository.Update(ctx, tx, r.UserID, r.StorySlug, r.ChapterSlug, c)
	if err != nil {
		return nil, err
	}

	story, err = cs.storyRepository.FindBySlugAndUserID(ctx, tx, r.StorySlug, r.UserID)
	if err != nil {
		return nil, err
	}

	chapter, err := cs.chapterRepository.FindByStorySlugAndChapterSlug(ctx, tx, r.UserID, story.Slug, updatedChapter.Slug)
	if err != nil {
		return nil, err
	}

	return &model.UpdatedChapterResponse{
		Id:            chapter.Id,
		StoryID:       chapter.StoryID,
		Title:         chapter.Title,
		Slug:          chapter.Slug,
		Body:          chapter.Body,
		AuthorComment: chapter.AuthorComment,
		WordCounts:    chapter.WordCounts,
		ReadingTime:   chapter.ReadingTime,
		IsPublished:   chapter.IsPublished,
		CreatedAt:     time.UnixToTime(chapter.CreatedAt),
		UpdatedAt:     time.UnixToTime(chapter.UpdatedAt),
	}, nil
}

func (cs *chapterService) GetChapterByStorySlugAndChapterSlug(ctx context.Context, userID uint64, storySlug string, chapterSlug string) (*model.ChapterResponseByStorySlugAndChapterSlug, error) {
	tx, err := cs.db.Begin()
	if err != nil {
		return nil, err
	}
	defer database.CommitOrRollback(tx)

	story, err := cs.storyRepository.FindBySlugAndUserID(ctx, tx, storySlug, userID)
	if err != nil {
		return nil, err
	}

	chapter, err := cs.chapterRepository.FindByStorySlugAndChapterSlug(ctx, tx, userID, storySlug, chapterSlug)
	if err != nil {
		return nil, err
	}

	return &model.ChapterResponseByStorySlugAndChapterSlug{
		Id:      chapter.Id,
		StoryID: story.Id,
		Story: storymodel.StoryResponseByChapter{
			Id:    chapter.Story.Id,
			Slug:  chapter.Story.Slug,
			Title: chapter.Story.Title,
		},
		User: usermodel.UserResponseByChapter{
			Id:       chapter.User.Id,
			Name:     chapter.User.Name,
			Username: chapter.User.Username,
		},
		Title:         chapter.Title,
		Slug:          chapter.Slug,
		Body:          chapter.Body,
		AuthorComment: chapter.AuthorComment,
		ReadingTime:   chapter.ReadingTime,
		UpdatedAt:     time.UnixToTime(chapter.UpdatedAt),
	}, nil
}

func (cs *chapterService) RemoveChapter(ctx context.Context, userID uint64, chapterID string) (*model.DeletedChapterResponse, error) {
	tx, err := cs.db.Begin()
	if err != nil {
		return nil, err
	}
	defer database.CommitOrRollback(tx)

	cId, err := strconv.Atoi(chapterID)

	if err != nil {
		return nil, err
	}

	_, err = cs.chapterRepository.FindByID(ctx, tx, uint64(cId))
	if err != nil {
		return nil, err
	}

	err = cs.chapterRepository.Delete(ctx, tx, userID, uint64(cId))
	if err != nil {
		return nil, err
	}

	return &model.DeletedChapterResponse{
		Status: true,
	}, nil
}

func (cs *chapterService) CalculateReadingTimeByChapters(ctx context.Context, chapters []entity.Chapter) (string, error) {
	var wordCountTotal uint64

	for _, c := range chapters {
		wordCountTotal += c.WordCounts
	}

	result := (wordCountTotal / 200)

	return fmt.Sprintf("%s Minutes", strconv.Itoa(int(result))), nil
}

func (cs *chapterService) GetAllChapterByStorySlug(ctx context.Context, storySlug string) (*[]model.ChapterResponseBySlug, error) {
	tx, err := cs.db.Begin()
	if err != nil {
		return nil, err
	}
	defer database.CommitOrRollback(tx)

	chapters, err := cs.chapterRepository.FindAllChapterByStorySlug(ctx, tx, storySlug)
	if err != nil {
		return nil, err
	}

	var chaptersResponse []model.ChapterResponseBySlug
	for _, c := range *chapters {
		chapterResponse := model.ChapterResponseBySlug{
			Id:         c.Id,
			StoryID:    c.StoryID,
			Title:      c.Title,
			Slug:       c.Slug,
			WordCounts: c.WordCounts,
		}

		chaptersResponse = append(chaptersResponse, chapterResponse)
	}

	return &chaptersResponse, nil
}

func (cs *chapterService) GetAllChapterByStoryID(ctx context.Context, storyID string) (*[]model.ChapterResponseBySlug, error) {
	tx, err := cs.db.Begin()
	if err != nil {
		return nil, err
	}
	defer database.CommitOrRollback(tx)

	sId, err := strconv.Atoi(storyID)
	if err != nil {
		return nil, err
	}

	_, err = cs.storyRepository.FindByID(ctx, tx, uint64(sId))
	if err != nil {
		return nil, err
	}

	chapters, err := cs.chapterRepository.FindAllChapterByStoryID(ctx, tx, uint64(sId))
	if err != nil {
		return nil, err
	}

	var chaptersResponse []model.ChapterResponseBySlug
	for _, c := range *chapters {
		chapterLikes, err := cs.chapterRepository.CountChapterLikesByChapterID(ctx, tx, c.Id)
		if err != nil {
			return nil, err
		}

		chapterResponse := model.ChapterResponseBySlug{
			Id:         c.Id,
			StoryID:    c.StoryID,
			Title:      c.Title,
			Slug:       c.Slug,
			WordCounts: c.WordCounts,
			Likes:      *chapterLikes,
		}

		chaptersResponse = append(chaptersResponse, chapterResponse)
	}

	return &chaptersResponse, nil
}

func NewService(cr chapter.ChapterRepository, sr story.StoryRepository, db *sql.DB) ChapterService {
	return &chapterService{
		chapterRepository: cr,
		storyRepository:   sr,
		db:                db,
	}
}
