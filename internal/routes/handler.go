package routes

import (
	"github.com/calebtracey/mind-your-business-api/internal/routes/endpoints"
	"github.com/go-chi/chi/v5"
	"github.com/swaggo/http-swagger"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
)

type Handler struct {
	Router endpoints.RouterI
}

const (
	basePath = "/api/v1"
	Path     = "doc.json"
)

func (h Handler) Routes() *chi.Mux {
	r := chi.NewRouter()
	setMiddleware(r)

	r.Get("/health", h.Router.Health())

	r.Route(basePath, func(r chi.Router) {
		r.Post("/newUser", h.Router.NewUser())
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:6080/swagger/"+Path), //The url pointing to API definition
	))

	return r
}
