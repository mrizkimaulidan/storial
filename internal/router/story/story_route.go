package story

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	storyhandler "github.com/mrizkimaulidan/storial/internal/handler/story"
	"github.com/mrizkimaulidan/storial/internal/middleware"
	chapterrepo "github.com/mrizkimaulidan/storial/internal/repository/chapter"
	storyrepo "github.com/mrizkimaulidan/storial/internal/repository/story"
	chapterservice "github.com/mrizkimaulidan/storial/internal/service/chapter"
	storyservice "github.com/mrizkimaulidan/storial/internal/service/story"
)

// Register routes.
func RegisterRoutes(r *mux.Router, db *sql.DB) {
	storyRepository := storyrepo.NewRepository()
	chapterRepository := chapterrepo.NewRepository()
	chapterService := chapterservice.NewService(chapterRepository, storyRepository, db)
	storyService := storyservice.NewService(storyRepository, chapterRepository, chapterService, db)
	storyHandler := storyhandler.NewHandler(storyService)

	v1 := r.PathPrefix("/api/v1").Subrouter()

	middleware := middleware.New()

	v1.Handle("/add-book", storyHandler.Store()).Methods(http.MethodPost)
	v1.Handle("/edit-book/{slug}", storyHandler.Update()).Methods(http.MethodPut, http.MethodPatch)
	v1.Handle("/book_front/{filename}", storyHandler.LoadImageCover()).Methods(http.MethodGet)
	v1.Handle("/writers/book/{id}/delete", storyHandler.Delete()).Methods(http.MethodDelete)
	v1.Handle("/user/books", storyHandler.GetAll()).Methods(http.MethodGet)
	v1.Handle("/book/{slug}", storyHandler.GetBySlug()).Methods(http.MethodGet)
	v1.Handle("/book-list", storyHandler.Filter()).Methods(http.MethodGet)
	v1.Handle("/{categorySlug}", storyHandler.FilterByCategorySlug()).Methods(http.MethodGet)
	v1.Use(middleware.JWTAuthorization)
}
