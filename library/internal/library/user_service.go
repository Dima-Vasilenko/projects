package library

import (
	"library/internal/errors_alias"
	"library/internal/user"
	"strings"
)

type UserRepository interface {
	AddUser(u *user.User) error
	GetUserByUsername(us string) (user.User, bool)
	GetListOfUsers() []user.User
	IsExist(us string) bool
	IsIDExist(id int) bool
}

type UserService struct {
	usRepo UserRepository
}

func NewUserService(usMem UserRepository) *UserService {
	return &UserService{
		usRepo: usMem,
	}
}

func (u *UserService) RegisterUser(us string) (int, error) {
	us = strings.TrimSpace(us)
	if us == "" {
		return -1, errors_alias.ErrIncorrectUserData
	}

	if exists := u.usRepo.IsExist(us); exists {
		return -1, errors_alias.ErrRepeatingUsername
	}

	uAdd := user.NewUser(us)
	if err := u.usRepo.AddUser(uAdd); err != nil {
		return -1, err
	}
	return uAdd.ID, nil
}

func (u *UserService) LogIn(us string) (int, error) {
	user, ok := u.usRepo.GetUserByUsername(us)
	if !ok {
		return -1, errors_alias.ErrUnknownUser
	}
	return user.ID, nil
}

func (u *UserService) GetListOfUsers() []user.User {
	return u.usRepo.GetListOfUsers()
}
