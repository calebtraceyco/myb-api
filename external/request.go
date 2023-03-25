package external

import "github.com/calebtraceyco/mind-your-business-api/external/models"

type ApiRequest struct {
	Payload struct {
		Request  Request `json:"request,omitempty"`
		Endpoint string  `json:"endpoint,omitempty"`
	} `json:"payload"`
}

type Request struct {
	User *models.User `json:"user,omitempty"`
}

func (r *Request) UserRequest(user *models.User) {
	r.User = user
}
