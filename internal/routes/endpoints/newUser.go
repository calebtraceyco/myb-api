package endpoints

import (
	"encoding/json"
	"github.com/calebtraceyco/mind-your-business-api/external"
	"github.com/calebtraceyco/mind-your-business-api/external/endpoints"
	"github.com/calebtraceyco/mind-your-business-api/internal/facade"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// NewUserHandler route handler for /newUser endpoint
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
func (r Router) NewUserHandler(service facade.ServiceI) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {

		rw.Header().Set("Content-Type", "application/json")
		var apiRequest external.ApiRequest
		var apiResponse *external.Response

		if err := json.NewDecoder(req.Body).Decode(&apiRequest); err != nil {
			log.Errorf("NewUserHandler: decode error: %v", err)
			rw.WriteHeader(400)
			return
		}

		apiRequest.Payload.Endpoint = endpoints.NewUser

		if apiResponse = service.UserResponse(req.Context(), apiRequest); len(apiResponse.Message.ErrorLog) > 0 {
			log.Errorf("NewUserHandler: \"/newUser\" - %v", apiResponse.Message.ErrorLog)
			rw.WriteHeader(500)

		} else {
			bytes, _ := json.Marshal(apiResponse)
			_, _ = rw.Write(bytes)
			rw.WriteHeader(200)
		}
	}
}
