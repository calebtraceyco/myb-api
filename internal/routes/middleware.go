package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"time"
)

func setMiddleware(r *chi.Mux) {
	// Injects a request ID into the context of each request
	r.Use(middleware.RequestID)
	// Sets a http.Request's RemoteAddr to either X-Real-IP or X-Forwarded-For
	r.Use(middleware.RealIP)
	// Logs the start and end of each request with the elapsed processing time
	r.Use(middleware.Logger)
	// Gracefully absorb panics and prints the stack trace
	r.Use(middleware.Recoverer)
	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{LocalhostSwagger, LocalhostCRA, LocalhostVite, LocalhostVite2, GithubPages, GithubPages1, GithubPages2},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodOptions, http.MethodDelete, http.MethodPut},
		AllowedHeaders:   []string{"Access-Control-Allow-Methods", "Access-Control-Allow-Origin", "X-Requested-With", "Authorization", "Content-Type", "X-Requested-With", "Bearer", "Origin"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
}

const (
	LocalhostSwagger = "http://localhost:6080/swagger/"
	LocalhostCRA     = "http://localhost:3000"
	LocalhostVite    = "http://localhost:5173"
	LocalhostVite2   = "http://localhost:5173/robot-image-ui/"
	GithubPages      = "https://calebtracey.github.io/robot-image-ui"
	GithubPages1     = "https://calebtracey.github.io"
	GithubPages2     = "https://calebtracey.github.io/robot-image-ui/"
)
