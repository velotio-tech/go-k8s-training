package order

type Order struct {
	OrderId string   `bson:"orderId" json:"orderId"`
	UserId  string   `bson:"userId" json:"userId"`
	Items   []string `bson:"items" json:"items"`
}

type AddItemDTO struct {
	OrderId string `json:"orderId" binding:"required"`
	UserId  string `json:"userId" binding:"required"`
	Item    string `json:"itemId" binding:"required"`
}

type OrderAllDTO struct {
	UserId string `json:"userId" binding:"required"`
}

type OrderDTO struct {
	OrderId string `json:"orderId" binding:"required"`
	UserId  string `json:"userId" binding:"required"`
}
