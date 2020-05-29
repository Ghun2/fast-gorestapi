package user

import (
	"github.com/Ghun2/fast-gorestapi/model"
)

type Store interface {
	Create(*model.User) error
	GetByAuthID(string) (*model.User, error)
	GetByID(uint) (*model.User, error)
	Update(*model.User) error
	Delete(*model.User) error
	ListUsers() ([]model.User, error)
}