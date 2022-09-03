package chapter

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mrizkimaulidan/storial/internal/model/story"
	"github.com/mrizkimaulidan/storial/internal/model/user"
)

type CreateChapterRequest struct {
	UserID        uint64
	StorySlug     string
	Title         string
	Body          string
	AuthorComment string
	IsPublished   string
}

func (ccr *CreateChapterRequest) Validate() error {
	return validation.ValidateStruct(ccr,
		validation.Field(&ccr.UserID, validation.Required),
		validation.Field(&ccr.StorySlug, validation.Required, validation.Length(5, 255)),
		validation.Field(&ccr.Title, validation.Required, validation.Length(5, 255)),
		validation.Field(&ccr.Body, validation.Required, validation.Length(5, 4294967295)),
		validation.Field(&ccr.AuthorComment, validation.Length(0, 255)),
		validation.Field(&ccr.IsPublished, validation.Required, validation.In("0", "1")),
	)
}

type CreatedChapterdResponse struct {
	Id            uint64    `json:"id"`
	StoryID       uint64    `json:"storyId"`
	Title         string    `json:"title"`
	Slug          string    `json:"slug"`
	Body          string    `json:"body"`
	AuthorComment string    `json:"authorComment"`
	WordCounts    uint64    `json:"wordCounts"`
	Likes         uint64    `json:"likes"`
	ReadingTime   string    `json:"readingTime"`
	IsPublished   bool      `json:"isPublished"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type UpdateChapterRequest struct {
	UserID        uint64
	StorySlug     string
	ChapterSlug   string
	Title         string
	Body          string
	AuthorComment string
	IsPublished   string
}

func (ucr *UpdateChapterRequest) Validate() error {
	return validation.ValidateStruct(ucr,
		validation.Field(&ucr.UserID, validation.Required),
		validation.Field(&ucr.StorySlug, validation.Required),
		validation.Field(&ucr.ChapterSlug, validation.Required),
		validation.Field(&ucr.Title, validation.Required, validation.Length(5, 255)),
		validation.Field(&ucr.Body, validation.Required, validation.Length(5, 4294967295)),
		validation.Field(&ucr.AuthorComment, validation.Length(0, 255)),
		validation.Field(&ucr.IsPublished, validation.In("0", "1")),
	)
}

type UpdatedChapterResponse struct {
	Id            uint64    `json:"id"`
	StoryID       uint64    `json:"storyId"`
	Title         string    `json:"title"`
	Slug          string    `json:"slug"`
	Body          string    `json:"body"`
	AuthorComment string    `json:"authorComment"`
	WordCounts    uint64    `json:"wordCounts"`
	Likes         uint64    `json:"likes"`
	ReadingTime   string    `json:"readingTime"`
	IsPublished   bool      `json:"isPublished"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type ChapterResponse struct {
	Id            uint64                       `json:"id"`
	StoryID       uint64                       `json:"storyId"`
	Story         story.StoryResponseByChapter `json:"story"`
	User          user.UserResponseByChapter   `json:"user"`
	Title         string                       `json:"title"`
	Slug          string                       `json:"slug"`
	Body          string                       `json:"body"`
	AuthorComment string                       `json:"authorComment"`
	WordCounts    uint64                       `json:"wordCounts"`
	Likes         uint64                       `json:"likes"`
	ReadingTime   string                       `json:"readingTime"`
	IsPublished   bool                         `json:"isPublished"`
	CreatedAt     time.Time                    `json:"createdAt"`
	UpdatedAt     time.Time                    `json:"updatedAt"`
}

type ChapterResponseByStorySlugAndChapterSlug struct {
	Id            uint64                       `json:"id"`
	StoryID       uint64                       `json:"storyId"`
	Story         story.StoryResponseByChapter `json:"story"`
	User          user.UserResponseByChapter   `json:"user"`
	Title         string                       `json:"title"`
	Slug          string                       `json:"slug"`
	Body          string                       `json:"body"`
	AuthorComment string                       `json:"authorComment"`
	Likes         uint64                       `json:"likes"`
	ReadingTime   string                       `json:"readingTime"`
	UpdatedAt     time.Time                    `json:"updatedAt"`
}

type ChapterResponseBySlug struct {
	Id         uint64 `json:"id"`
	StoryID    uint64 `json:"storyId"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	WordCounts uint64 `json:"wordCounts"`
	Likes      uint64 `json:"likes"`
}

type DeletedChapterResponse struct {
	Status bool `json:"status"`
}

type LikedChapterResponse struct {
	Status bool `json:"status"`
}
