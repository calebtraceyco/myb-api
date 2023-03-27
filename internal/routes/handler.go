package routes

import (
	routes "github.com/calebtraceyco/mind-your-business-api/external/endpoints"
	"github.com/calebtraceyco/mind-your-business-api/internal/facade"
	"github.com/calebtraceyco/mind-your-business-api/internal/routes/endpoints"
	"github.com/go-chi/chi/v5"
	"github.com/swaggo/http-swagger"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
)

type Handler struct {
	Router endpoints.RouterI
}

func (h Handler) RouteHandler(service facade.ServiceI) *chi.Mux {
	r := chi.NewRouter()
	setMiddleware(r)

	r.Get(routes.Health, h.Router.Health())

	r.Route(v1BasePath, func(r chi.Router) {
		r.Post(routes.NewUser, h.Router.NewUserHandler(service))
	})

	// serve swagger static page: http://localhost:6080/swagger/index.html
	r.Route(swaggerBasePath, func(r chi.Router) {
		r.Get(wildCard, httpSwagger.Handler(
			httpSwagger.URL(swaggerUiPath+swaggerDoc)), //The url pointing to API definition
		)
	})

	return r
}

const (
	wildCard        = "/*"
	v1BasePath      = "/api/v1"
	swaggerBasePath = "/swagger"
	swaggerDoc      = "doc.json"
	swaggerUiPath   = "http://localhost:6080/swagger/"
)
