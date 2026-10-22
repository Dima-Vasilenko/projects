package user

import (
	"fmt"
)

type User struct {
	Username string
	ID       int
}

func (u *User) String() string {
	return fmt.Sprintf("User: %s", u.Username)
}
