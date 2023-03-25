package endpoints

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// Health route handler for /health endpoint
//
// @Summary      Health check endpoint
// @Description  request to check for 200 response
// @Tags         util
// @Accept       json
// @Produce      json
// @Success      200  {object}  http.HandlerFunc
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /health [post]
func (r *Router) Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Traceln("=== Health Check === \n" + time.Now().Local().String())

		w.WriteHeader(http.StatusOK)
	}
}
