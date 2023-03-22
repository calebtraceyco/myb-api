package routes

import (
	"github.com/calebtracey/mind-your-business-api/internal/facade"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Handler struct {
	Service facade.ServiceI
}

func (h Handler) Routes() *chi.Mux {
	r := chi.NewRouter()

	setMiddleware(r)

	r.Post("/newUser", h.newUser())

	return r
}

func (h Handler) newUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if _, err := h.Service.NewUser(r.Context(), r.GetBody); err != nil {
			log.Panicf("/newUser - %v", err)
			w.WriteHeader(500)

		} else {
			_, _ = w.Write([]byte("{message: user added successfully}"))
			w.WriteHeader(200)

		}

	}
}

func setMiddleware(r *chi.Mux) {
	// docs: https://github.com/go-chi/chi
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
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: allowedOrigins,
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: allowedMethods,
		AllowedHeaders: allowedHeaders,
		//ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
}

var (
	allowedOrigins = []string{localhostCRA, localhostVite, localhostVite2, githubPages, githubPages1, githubPages2}
	allowedMethods = []string{"GET", "POST", "OPTIONS", "DELETE", "PUT"}
	allowedHeaders = []string{"Access-Control-Allow-Methods", "Access-Control-Allow-Origin", "X-Requested-With", "Authorization", "Content-Type", "X-Requested-With", "Bearer", "Origin"}
)

const (
	localhostCRA   = "http://localhost:3000"
	localhostVite  = "http://localhost:5173"
	localhostVite2 = "http://localhost:5173/robot-image-ui/"
	githubPages    = "https://calebtracey.github.io/robot-image-ui"
	githubPages1   = "https://calebtracey.github.io"
	githubPages2   = "https://calebtracey.github.io/robot-image-ui/"
)
