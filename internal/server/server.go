package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/mrizkimaulidan/storial/internal/config"
	"github.com/mrizkimaulidan/storial/internal/database"
	"github.com/mrizkimaulidan/storial/internal/middleware"
	"github.com/mrizkimaulidan/storial/internal/router/authentication"
	"github.com/mrizkimaulidan/storial/internal/router/category"
	"github.com/mrizkimaulidan/storial/internal/router/chapter"
	"github.com/mrizkimaulidan/storial/internal/router/story"
)

type Server struct {
	router *mux.Router
	c      *config.Config
}

func NewServer() *Server {
	return &Server{
		router: mux.NewRouter(),
		c:      config.New().GetConfig(),
	}
}

// Run the server asynchronous using goroutine.
// The server will be closed until cancel signal received.
// And handling the shutdown gracefully.
func (s *Server) Run() {
	s.routes()

	middleware := middleware.New()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", s.c.APP_PORT),
		Handler: middleware.LoggingMiddleware(s.router),
	}

	go func() {
		log.Println("server running at", srv.Addr)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalln("error running the server", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan
	log.Println("got signal", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	log.Println("shutting down server..")
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatalln("error shutting down server", err)
	}

	log.Println("server shutted down successfully")
}

// Setup routes endpoint.
func (s *Server) routes() {
	db := database.NewDatabase().Open()

	authentication.RegisterRoutes(s.router, db)
	story.RegisterRoutes(s.router, db)
	chapter.RegisterRoutes(s.router, db)
	category.RegisterRoutes(s.router, db)
}
