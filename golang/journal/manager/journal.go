package manager

import "fmt"

type Journal struct {
	Id        int
	CreatedAt string
	Message   string
}

type User struct {
	Name         string
	Email        string
	Password     string
	IsRegistered bool
	IsLoggedIn   bool
	Journal      []Journal
}

func AddEntry(message string) {
	fmt.Println(message)
}

func ListEntry(email string) {
	fmt.Println(email)
}

func PrintStruct() {
	journal1 := Journal{Id: 1, CreatedAt: "today", Message: "Hii"}
	journal2 := Journal{Id: 2, CreatedAt: "today", Message: "Hello"}
	tmp := []Journal{}
	tmp = append(tmp, journal1, journal2)
	alex := User{Name: "Alex", Email: "alex@email.com", Password: "password", IsRegistered: true, IsLoggedIn: false, Journal: tmp}

	fmt.Println(alex)
}
