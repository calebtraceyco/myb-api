package routes

import (
	"github.com/calebtracey/mind-your-business-api/internal/facade"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type Handler struct {
	Service facade.ServiceI
}

func (h Handler) Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", handler())
	return r
}

func handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		_, _ = w.Write([]byte("{status:ok}"))
		w.WriteHeader(200)

	}
}
