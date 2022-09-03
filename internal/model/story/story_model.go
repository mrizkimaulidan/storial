package story

import (
	"mime/multipart"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	categorymodel "github.com/mrizkimaulidan/storial/internal/model/category"
	usermodel "github.com/mrizkimaulidan/storial/internal/model/user"
)

type CreateStoryRequest struct {
	UserID          uint64
	CategoryID      string
	Title           string
	Description     string
	Cover           multipart.File
	CoverFileheader *multipart.FileHeader
	IsAdult         string
	IsPublished     string
}

func (cr *CreateStoryRequest) Validate() error {
	return validation.ValidateStruct(cr,
		validation.Field(&cr.UserID, validation.Required),
		validation.Field(&cr.CategoryID, validation.Required),
		validation.Field(&cr.Title, validation.Required, validation.Length(5, 255)),
		validation.Field(&cr.Description, validation.Required, validation.Length(5, 16777215)),
		validation.Field(&cr.IsAdult, validation.Required, validation.In("0", "1")),
		validation.Field(&cr.IsPublished, validation.Required, validation.In("0", "1")),
	)
}

type CreatedStoryResponse struct {
	Id          uint64                         `json:"id"`
	UserID      uint64                         `json:"userId"`
	CategoryID  uint64                         `json:"categoryId"`
	Category    categorymodel.CategoryResponse `json:"category"`
	Title       string                         `json:"title"`
	Slug        string                         `json:"slug"`
	Description string                         `json:"description"`
	IsAdult     bool                           `json:"isAdult"`
	IsPublished bool                           `json:"isPublished"`
	Cover       string                         `json:"cover"`
	CreatedAt   time.Time                      `json:"createdAt"`
	UpdatedAt   time.Time                      `json:"updatedAt"`
}

type UpdateStoryRequest struct {
	Slug            string
	UserID          uint64
	CategoryID      string
	Title           string
	Description     string
	Cover           multipart.File
	CoverFileheader *multipart.FileHeader
	IsAdult         string
	IsPublished     string
}

func (usr *UpdateStoryRequest) Validate() error {
	return validation.ValidateStruct(usr,
		validation.Field(&usr.UserID, validation.Required),
		validation.Field(&usr.CategoryID, validation.Required),
		validation.Field(&usr.Title, validation.Required, validation.Length(5, 255)),
		validation.Field(&usr.Description, validation.Required, validation.Length(5, 16777215)),
		validation.Field(&usr.IsAdult, validation.Required, validation.In("0", "1")),
		validation.Field(&usr.IsPublished, validation.Required, validation.In("0", "1")),
	)
}

type UpdatedStoryResponse struct {
	Id          uint64                         `json:"id"`
	UserID      uint64                         `json:"userId"`
	CategoryID  uint64                         `json:"categoryId"`
	Category    categorymodel.CategoryResponse `json:"category"`
	Title       string                         `json:"title"`
	Slug        string                         `json:"slug"`
	Description string                         `json:"description"`
	IsAdult     bool                           `json:"isAdult"`
	IsPublished bool                           `json:"isPublished"`
	Cover       string                         `json:"cover"`
	CreatedAt   time.Time                      `json:"createdAt"`
	UpdatedAt   time.Time                      `json:"updatedAt"`
}

type DeleteStoryRequest struct {
	Id string
}

type DeletetedStoryResponse struct {
	Status bool `json:"status"`
}

type StoryResponse struct {
	Id          uint64                 `json:"id"`
	UserID      uint64                 `json:"userId"`
	User        usermodel.UserResponse `json:"user"`
	Title       string                 `json:"title"`
	Slug        string                 `json:"slug"`
	IsPublished bool                   `json:"isPublished"`
	CreatedAt   time.Time              `json:"createdAt"`
	UpdatedAt   time.Time              `json:"updatedAt"`
}

type StoryResponseByChapter struct {
	Id    uint64 `json:"id"`
	Title string `json:"title"`
	Slug  string `json:"slug"`
}

type StoryResponseBySlug struct {
	Id            uint64                                    `json:"id"`
	UserID        uint64                                    `json:"userId"`
	User          usermodel.UserResponseBySlug              `json:"user"`
	CategoryID    uint64                                    `json:"categoryId"`
	Category      categorymodel.CategoryResponseByStorySlug `json:"category"`
	Title         string                                    `json:"title"`
	Slug          string                                    `json:"slug"`
	ChapterCounts uint64                                    `json:"chapterCounts"`
	ReadingTime   string                                    `json:"readingTime"`
	CreatedAt     time.Time                                 `json:"createdAt"`
	UpdatedAt     time.Time                                 `json:"updatedAt"`
}

type StoryResponseByFilter struct {
	Id            uint64                         `json:"id"`
	UserID        uint64                         `json:"userId"`
	User          usermodel.UserResponseByFilter `json:"user"`
	Title         string                         `json:"title"`
	Slug          string                         `json:"slug"`
	ChapterCounts uint64                         `json:"chapterCounts"`
	IsAdult       bool                           `json:"isAdult"`
	Cover         string                         `json:"cover"`
	CreatedAt     time.Time                      `json:"createdAt"`
}

type StoryResponseByCategorySlug struct {
	Id            uint64                       `json:"id"`
	UserID        uint64                       `json:"userId"`
	User          usermodel.UserResponseBySlug `json:"user"`
	Title         string                       `json:"title"`
	Slug          string                       `json:"slug"`
	ChapterCounts uint64                       `json:"chapterCounts"`
	CreatedAt     time.Time                    `json:"createdAt"`
	UpdatedAt     time.Time                    `json:"updatedAt"`
}
