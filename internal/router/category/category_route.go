package category

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	categoryhandler "github.com/mrizkimaulidan/storial/internal/handler/category"
	"github.com/mrizkimaulidan/storial/internal/middleware"
	categoryrepo "github.com/mrizkimaulidan/storial/internal/repository/category"
	storyrepo "github.com/mrizkimaulidan/storial/internal/repository/story"
	categoryservice "github.com/mrizkimaulidan/storial/internal/service/category"
)

// Register routes.
func RegisterRoutes(r *mux.Router, db *sql.DB) {
	categoryRepository := categoryrepo.NewRepository()
	storyRepository := storyrepo.NewRepository()
	categoryService := categoryservice.NewService(categoryRepository, storyRepository, db)
	categoryHandler := categoryhandler.NewHandler(categoryService)

	middleware := middleware.New()

	v1 := r.PathPrefix("/api/v1").Subrouter()
	v1.Handle("/books/categories", categoryHandler.GetAllCategory()).Methods(http.MethodGet)
	v1.Use(middleware.JWTAuthorization)
}
