package endpoints

import (
	"github.com/calebtracey/mind-your-business-api/internal/facade"
	"net/http"
)

type RouterI interface {
	NewUser() http.HandlerFunc
	Health() http.HandlerFunc
}

type Router struct {
	Service facade.ServiceI
}

const (
	Health  = "/health"
	NewUser = "/newUser"
)
