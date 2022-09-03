package chapter

import "net/http"

type ChapterHandler interface {
	AddChapter() http.Handler
	EditChapter() http.Handler
	GetChapter() http.Handler
	DeleteChapter() http.Handler
	GetChapters() http.Handler
	LikeChapter() http.Handler
}
