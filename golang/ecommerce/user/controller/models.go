package controller

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func init() {
	db := Connect()
	db.AutoMigrate(&User{})
}
