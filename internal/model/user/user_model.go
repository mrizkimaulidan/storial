package user

import "time"

type UserResponse struct {
	Id    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserResponseByChapter struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type UserResponseBySlug struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type UserResponseByFilter struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type UserResponseByLatest struct {
	Id            uint64    `json:"id"`
	Name          string    `json:"name"`
	Username      string    `json:"username"`
	ChapterCounts uint64    `json:"chapterCounts"`
	CreatedAt     time.Time `json:"createdAt"`
}
