package store

import "github.com/Xlussov/Rest-api/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}