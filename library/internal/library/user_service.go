package library

import (
	"errors"
	"fmt"
	"library/internal/user"
	"slices"
)

type UserRepository interface {
	AddUser(u *user.User)
	GetUserByUsername(us string) (*user.User, bool)
	GetListOfUsers() []*user.User
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
	if us == "" {
		return -1, errors.New("incorrect input")
	}
	exists := slices.ContainsFunc(u.usRepo.GetListOfUsers(), func(u *user.User) bool {
		return u.Username == us
	})
	if exists {
		return -1, ErrRepeatingUsername
	}

	uAdd := user.NewUser(us)
	u.usRepo.AddUser(uAdd)
	return uAdd.ID, nil
}

func (u *UserService) LogIn(us string) (int, error) {
	index := slices.IndexFunc(u.usRepo.GetListOfUsers(), func(u *user.User) bool {
		return u.Username == us
	})
	if index == -1 {
		return -1, ErrUncknownUser
	}
	return u.usRepo.GetListOfUsers()[index].ID, nil
}

func (u *UserService) ShowListOfUsers() {
	for _, val := range u.usRepo.GetListOfUsers() {
		fmt.Println(val)
	}
}
