package user

import (
	"fmt"
)

type User struct {
	Username string
	Id       int
}

func (u *User) String() string {
	return fmt.Sprintf("User: %s\n", u.Username)
}
