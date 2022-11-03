package cmd

type Manager struct {
	registeredUsers map[string]*User
	loggedInUsers   map[string]*User
}
