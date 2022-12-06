package manager

type Journal struct {
	CreatedAt string `json:"createdAt"`
	Message   string `json:"message"`
}

type User struct {
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	IsLoggedIn bool      `json:"isloggedin"`
	Journal    []Journal `json:"journal"`
}

type users []User

func NewUser(name, email, password string) User {
	// &variable = address
	// *address = variable
	u := User{Name: name, Email: email, Password: password, IsLoggedIn: false, Journal: make([]Journal, 0)}
	return u
}
