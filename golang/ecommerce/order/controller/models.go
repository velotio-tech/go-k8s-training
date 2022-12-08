package controller

type Order struct {
	Id     int    `json:"id"`
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
}

func init() {
	db := Connect()
	db.AutoMigrate(&Order{})
}
