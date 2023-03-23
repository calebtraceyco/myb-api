package external

import "github.com/calebtraceyco/mind-your-business-api/external/models"

type ApiRequest struct {
	Request Request `json:"request,omitempty"`
}

type Request struct {
	User *models.User `json:"user,omitempty"`
}

func (r *Request) UserRequest(user *models.User) {
	r.User = user
}
