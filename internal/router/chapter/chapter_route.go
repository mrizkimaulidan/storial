package chapter

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	chapterhandler "github.com/mrizkimaulidan/storial/internal/handler/chapter"
	"github.com/mrizkimaulidan/storial/internal/middleware"
	chapterrepository "github.com/mrizkimaulidan/storial/internal/repository/chapter"
	"github.com/mrizkimaulidan/storial/internal/repository/story"
	chapterservice "github.com/mrizkimaulidan/storial/internal/service/chapter"
)

// Register routes.
func RegisterRoutes(r *mux.Router, db *sql.DB) {
	chapterRepository := chapterrepository.NewRepository()
	storyRepository := story.NewRepository()
	chapterService := chapterservice.NewService(chapterRepository, storyRepository, db)
	chapterHandler := chapterhandler.NewHandler(chapterService)

	middleware := middleware.New()

	v1 := r.PathPrefix("/api/v1").Subrouter()
	v1.Handle("/add-chapter/{storySlug}", chapterHandler.AddChapter()).Methods(http.MethodPost)
	v1.Handle("/edit-chapter/{storySlug}/{chapterSlug}", chapterHandler.EditChapter()).Methods(http.MethodPut, http.MethodPatch)
	v1.Handle("/book/{storySlug}/{chapterSlug}", chapterHandler.GetChapter()).Methods(http.MethodGet)
	v1.Handle("/writers/chapter/{chapterId}/delete", chapterHandler.DeleteChapter()).Methods(http.MethodDelete)
	v1.Handle("/books/{storyId}/chapters", chapterHandler.GetChapters()).Methods(http.MethodGet)
	v1.Handle("/books/{storyId}/chapters/{chapterId}/votes/up", chapterHandler.LikeChapter()).Methods(http.MethodPost)
	v1.Use(middleware.JWTAuthorization)
}
