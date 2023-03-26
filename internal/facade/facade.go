package facade

import (
	"context"
	"github.com/calebtraceyco/mind-your-business-api/external"
	"github.com/calebtraceyco/mind-your-business-api/external/endpoints"
	"github.com/calebtraceyco/mind-your-business-api/internal/dao/user"
	log "github.com/sirupsen/logrus"
)

type ServiceI interface {
	UserResponse(ctx context.Context, apiRequest external.ApiRequest) (resp *external.Response)
}

type Service struct {
	UserDAO user.DAOI
}

func (s Service) UserResponse(ctx context.Context, apiRequest external.ApiRequest) (resp *external.Response) {
	resp = new(external.Response)
	// TODO validation
	if apiRequest.Payload.Request.User == nil {
		panic("missing user from payload")
	}

	switch apiRequest.Payload.Endpoint {
	case endpoints.NewUser:
		log.Traceln("UserResponse: /newUser endpoint")
		if daoResp, err := s.UserDAO.AddUser(ctx, apiRequest.Payload.Request.User); err != nil {
			resp.SetErrorLog([]error{err}, "UserResponse", "500")
			return resp
		} else {
			// TODO map user response
			resp.Details = []any{daoResp}
		}
	default:
		// TODO change to error and/or default call
		// panic for debugging
		panic("UserResponse: missing endpoint")
	}
	return resp
}
