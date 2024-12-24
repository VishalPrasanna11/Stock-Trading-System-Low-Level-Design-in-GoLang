package stockbroker

type UserList struct {
	users map[string]*User
}

func NewUserList() *UserList {
	return &UserList{users: make(map[string]*User)}
}

func (ul *UserList) GetUser(userId string) *User {
	return ul.users[userId]
}

func (ul *UserList) AddUser(userId string, user *User) {
	ul.users[userId] = user

}
