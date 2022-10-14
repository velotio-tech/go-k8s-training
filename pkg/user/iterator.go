package user

type userIterator struct {
	index int
	users []User
}

func createUserIterator(users []User) *userIterator {
	return &userIterator{
		users: users,
	}
}

func (u *userIterator) HasNext() bool {
	return u.index < len(u.users)
}

func (u *userIterator) Get() interface{} {
	if u.HasNext() {
		user := u.users[u.index]
		u.index++
		return &user
	}
	return nil
}
