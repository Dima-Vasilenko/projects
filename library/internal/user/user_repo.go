package user

func NewUser(us string) *User {
	user := &(User{
		Username: us,
	})
	return user
}

type UserMemory struct {
	users  []*User
	nextID int
}

func NewUserMemory() *UserMemory {
	return &(UserMemory{
		users:  make([]*User, 0),
		nextID: 0,
	})
}

func (um *UserMemory) AddUser(u *User) {
	u.ID = um.nextID
	um.nextID++
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
