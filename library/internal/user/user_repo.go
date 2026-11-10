package user

import (
	"library/internal/errors_alias"
	"slices"
)

func NewUser(us string) *User {
	return &(User{
		Username: us,
	})
}

type UserMemory struct {
	users  []User
	nextID int
}

func NewUserMemory() *UserMemory {
	return &(UserMemory{
		users:  make([]User, 0),
		nextID: 0,
	})
}

func (um *UserMemory) AddUser(u *User) error {
	if u == nil {
		return errors_alias.ErrUnknownUser
	}
	u.ID = um.nextID
	um.nextID++
	um.users = append(um.users, *u)
	return nil
}

func (um *UserMemory) GetUserByUsername(us string) (User, bool) {
	for _, val := range um.users {
		if val.Username == us {
			return val, true
		}
	}
	return User{}, false
}

func (um *UserMemory) GetListOfUsers() []User {
	return append([]User(nil), um.users...)
}

func (um *UserMemory) IsExist(us string) bool {
	exist := slices.ContainsFunc(um.users, func(u User) bool {
		return u.Username == us
	})
	return exist
}

func (um *UserMemory) IsIDExist(id int) bool {
	exist := slices.ContainsFunc(um.users, func(u User) bool {
		return u.ID == id
	})
	return exist
}
