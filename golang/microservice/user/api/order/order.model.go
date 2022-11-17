package order

type AddItemDTO struct {
	OrderId string `uri:"orderId" binding:"required"`
	UserId  string `uri:"userId" binding:"required"`
	Item    string `uri:"itemId" binding:"required"`
}

type OrderAllDTO struct {
	UserId string `uri:"userId" binding:"required"`
}

type OrderDTO struct {
	OrderId string `uri:"orderId" binding:"required"`
	UserId  string `uri:"userId" binding:"required"`
}

type OrderRequest struct {
	OrderId string `json:"orderId"`
	UserId  string `json:"userId"`
	Item    string `json:"itemId"`
}
