package endpoints

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

// NewUser route handler for /newUser endpoint
//
// @Summary      New User request
// @Description  request to add new user to the database
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  http.HandlerFunc
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /api/v1/newUser [post]
func (r *Router) NewUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		if _, err := r.Service.NewUser(req.Context(), req.GetBody); err != nil {
			log.Panicf("/newUser - %v", err)
			rw.WriteHeader(500)

		} else {
			_, _ = rw.Write([]byte("{message: user added successfully}"))
			rw.WriteHeader(200)
		}
	}
}
