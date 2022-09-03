package authentication

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	authenticationhandler "github.com/mrizkimaulidan/storial/internal/handler/authentication"
	authenticationrepo "github.com/mrizkimaulidan/storial/internal/repository/authentication"
	authenticationservice "github.com/mrizkimaulidan/storial/internal/service/authentication"
)

// Register routes.
func RegisterRoutes(r *mux.Router, db *sql.DB) {
	authenticationRepository := authenticationrepo.NewRepository()
	authenticationService := authenticationservice.NewService(authenticationRepository, db)
	authenticationhandler := authenticationhandler.NewHandler(authenticationService)

	v1 := r.PathPrefix("/api/v1").Subrouter()
	v1.Handle("/register", authenticationhandler.Register()).Methods(http.MethodPost)
	v1.Handle("/login", authenticationhandler.Login()).Methods(http.MethodPost)
}
