package routes

import (
	"github.com/calebtracey/mind-your-business-api/internal/routes/endpoints"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Router endpoints.RouterI
}

const basePath = "/api/v1"

func (h Handler) Routes() *chi.Mux {
	r := chi.NewRouter()

	setMiddleware(r)

	r.Get("/health", h.Router.Health())

	r.Route(basePath, func(r chi.Router) {
		r.Post("/newUser", h.Router.NewUser())
	})

	return r
}
