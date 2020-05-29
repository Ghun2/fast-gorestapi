package controller

import (
	"github.com/Ghun2/fast-gorestapi/service/user"
)

type Controller struct {
	userStore		user.Store
}

func NewController(us user.Store) *Controller {
	return &Controller{
		userStore: 		us,
	}
}