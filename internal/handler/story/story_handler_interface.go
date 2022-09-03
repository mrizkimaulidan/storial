package story

import "net/http"

type StoryHandler interface {
	GetAll() http.Handler
	Store() http.Handler
	Update() http.Handler
	LoadImageCover() http.Handler
	Delete() http.Handler
	GetBySlug() http.Handler
	FilterByCategorySlug() http.Handler
	Filter() http.Handler
}
