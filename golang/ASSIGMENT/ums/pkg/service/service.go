package service

// Service ...
type Service struct {
	OrderService OrderService
}

func NewService(orderService OrderService) *Service {
	return &Service{OrderService: orderService}
}
