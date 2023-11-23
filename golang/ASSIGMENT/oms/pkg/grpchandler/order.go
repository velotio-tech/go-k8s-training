package grpchandler

import (
	"context"
	"time"

	"oms/pkg/models"

	pb "github.myproto.com"
)

// CreateOrder ...
func (g *GrpcHandler) CreateOrder(_ context.Context, request *pb.Order) (*pb.OrderCreateResponse, error) {
	products := models.Products{}
	for _, p := range request.Products {
		products = append(products, models.Product{
			ID:       p.ProductID,
			Price:    float64(p.FinalPrice),
			Quantity: p.Quantity,
		})
	}
	orderID, err := g.domain.CreateOrder(&models.Order{
		BuyerUserID: request.BuyerUserID,
		Status:      request.Status,
		TotalAmount: request.TotalAmount,
		PaymentID:   request.PaymentID,
		Products:    products,
	})
	if err != nil {
		return nil, err
	}

	return &pb.OrderCreateResponse{
		ID: orderID,
	}, nil
}

// DeleteOrder ...
func (g *GrpcHandler) DeleteOrder(_ context.Context, request *pb.DeleteOrderByOrderIDRequest) (*pb.DeleteOrderByOrderIDResponse, error) {
	return &pb.DeleteOrderByOrderIDResponse{}, g.domain.DeleteOrderByOrderID(request.OrderID)
}

// DeleteOrders ...
func (g *GrpcHandler) DeleteOrders(_ context.Context, request *pb.DeleteOrdersByBuyerIDRequest) (*pb.DeleteOrdersByBuyerIDResponse, error) {
	return &pb.DeleteOrdersByBuyerIDResponse{}, g.domain.DeleteOrderByUserID(request.BuyerID)
}

// GetOrders ...
func (g *GrpcHandler) GetOrders(_ context.Context, request *pb.OrdersGetRequest) (*pb.OrdersGetResponse, error) {
	allOrders := []*pb.Order{}
	orders, err := g.domain.GetOrdersByUserID(request.BuyerUserID)
	if err != nil {
		return nil, err
	}
	for _, o := range orders {
		products := []*pb.Product{}
		for _, p := range o.Products {
			products = append(products, &pb.Product{
				ProductID:  p.ID,
				FinalPrice: float32(p.Price),
				Quantity:   p.Quantity,
			})
		}
		allOrders = append(allOrders, &pb.Order{
			ID:          o.ID,
			BuyerUserID: o.BuyerUserID,
			Status:      o.Status,
			TotalAmount: o.TotalAmount,
			PaymentID:   o.PaymentID,
			Products:    products,
			CreatedAt:   o.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   o.UpdatedAt.Format(time.RFC3339),
		})
	}

	return &pb.OrdersGetResponse{Orders: allOrders}, nil
}

// GetOrder ...
func (g *GrpcHandler) GetOrder(_ context.Context, request *pb.OrderGetRequest) (*pb.OrderGetResponse, error) {
	order, err := g.domain.GetOrder(request.OrderID)
	if err != nil {
		return nil, err
	}
	products := []*pb.Product{}
	for _, p := range order.Products {
		products = append(products, &pb.Product{
			ProductID:  p.ID,
			FinalPrice: float32(p.Price),
			Quantity:   p.Quantity,
		})
	}
	return &pb.OrderGetResponse{Order: &pb.Order{
		ID:          order.ID,
		BuyerUserID: order.BuyerUserID,
		Status:      order.Status,
		TotalAmount: order.TotalAmount,
		PaymentID:   order.PaymentID,
		Products:    products,
		CreatedAt:   order.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   order.UpdatedAt.Format(time.RFC3339),
	}}, nil
}
