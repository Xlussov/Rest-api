package teststore


import (
	"github.com/Xlussov/Rest-api/internal/app/store"
	"github.com/Xlussov/Rest-api/internal/app/model"

)

type Store struct {
	userRepository *UserRepository
}

func New() *Store {
	return &Store{}
}


func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.userRepository
}