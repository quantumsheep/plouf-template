package modules_users

import (
	"github.com/quantumsheep/plouf"
	"github.com/quantumsheep/plouf/example/entities"
	"github.com/quantumsheep/plouf/plouf_modules"
)

type UsersService struct {
	plouf.Service

	UserRepository *plouf_modules.Repository[entities.User]
}

func (s *UsersService) FindById(id int) (*entities.User, error) {
	return s.UserRepository.FindOne(id)
}

func (s *UsersService) CreateUser(user *entities.User) (*entities.User, error) {
	return user, s.UserRepository.Create(user)
}

func (s *UsersService) Find() ([]*entities.User, error) {
	return s.UserRepository.Find()
}
