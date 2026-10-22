package user

var userId int = 0

func NewUser(us string) *User {
	user := new(User{
		Username: us,
		Id:       userId,
	})
	userId++
	return user
}

type UserMemory struct {
	users []*User
}

func NewUserMemory() *UserMemory {
	return new(UserMemory{
		users: make([]*User, 0),
	})
}

func (um *UserMemory) AddUser(u *User) {
	um.users = append(um.users, u)
}

func (um *UserMemory) GetUserByUsername(us string) (*User, bool) {
	for _, val := range um.users {
		if val.Username == us {
			return val, true
		}
	}
	return nil, false
}

func (um *UserMemory) GetListOfUsers() []*User {
	return um.users
}
