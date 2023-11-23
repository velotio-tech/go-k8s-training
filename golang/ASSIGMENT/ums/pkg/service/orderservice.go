package service

import (
	"context"
	"net/http"
	"time"

	"ums/pkg/exception"
	"ums/pkg/models"

	pb "github.myproto.com"
)

// OrderService ...
type OrderService interface {
	CreateOrder(order *models.Order) (string, *exception.Exception)
	GetOrderByID(orderID string) (*models.Order, *exception.Exception)
	GetOrders(userID string) ([]models.Order, *exception.Exception)
	DeleteOrder(orderID string) *exception.Exception
	DeleteOrders(userID string) *exception.Exception
}

// OrderServiceClient ...
type OrderServiceClient struct {
	client pb.OrderManagementClient
}

// NewOrderServiceClient ...
func NewOrderServiceClient(client pb.OrderManagementClient) *OrderServiceClient {
	return &OrderServiceClient{client: client}
}

// CreateOrder ...
func (o *OrderServiceClient) CreateOrder(order *models.Order) (string, *exception.Exception) {
	products := []*pb.Product{}
	for _, p := range order.Products {
		products = append(products, &pb.Product{
			ProductID:  p.ID,
			FinalPrice: float32(p.Price),
			Quantity:   p.Quantity,
		})
	}
	res, err := o.client.CreateOrder(context.Background(), &pb.Order{
		BuyerUserID: order.BuyerUserID,
		Status:      order.Status,
		TotalAmount: order.TotalAmount,
		PaymentID:   order.PaymentID,
		Products:    products,
	})
	if err != nil {
		return "", &exception.Exception{
			Err:        err,
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
			StatusText: exception.STATUS_SERVICE_ERROR,
		}
	}

	return res.ID, nil
}

// GetOrderByID ...
func (o *OrderServiceClient) GetOrderByID(orderID string) (*models.Order, *exception.Exception) {
	products := models.Products{}
	res, err := o.client.GetOrder(context.Background(), &pb.OrderGetRequest{
		OrderID: orderID,
	})
	if err != nil {
		return nil, &exception.Exception{
			Err:        err,
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
			StatusText: exception.STATUS_SERVICE_ERROR,
		}
	}
	for _, p := range res.Order.Products {
		products = append(products, models.Product{
			ID:       p.ProductID,
			Price:    float64(p.FinalPrice),
			Quantity: p.Quantity,
		})
	}

	return &models.Order{
		ID:          res.Order.ID,
		BuyerUserID: res.Order.BuyerUserID,
		CreatedAt:   convertStringToTime(res.Order.CreatedAt),
		UpdatedAt:   convertStringToTime(res.Order.UpdatedAt),
		TotalAmount: res.Order.TotalAmount,
		PaymentID:   res.Order.PaymentID,
		Products:    products,
	}, nil
}

// GetOrders ...
func (o *OrderServiceClient) GetOrders(userID string) ([]models.Order, *exception.Exception) {
	orders := []models.Order{}
	res, err := o.client.GetOrders(context.Background(), &pb.OrdersGetRequest{
		BuyerUserID: userID,
	})
	if err != nil {
		return nil, &exception.Exception{
			Err:        err,
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
			StatusText: exception.STATUS_SERVICE_ERROR,
		}
	}
	for _, o := range res.Orders {
		products := models.Products{}
		for _, p := range o.Products {
			products = append(products, models.Product{
				ID:       p.ProductID,
				Price:    float64(p.FinalPrice),
				Quantity: p.Quantity,
			})
		}
		orders = append(orders, models.Order{
			ID:          o.ID,
			BuyerUserID: o.BuyerUserID,
			TotalAmount: o.TotalAmount,
			PaymentID:   o.PaymentID,
			CreatedAt:   convertStringToTime(o.CreatedAt),
			UpdatedAt:   convertStringToTime(o.UpdatedAt),
			Products:    products,
		})
	}

	return orders, nil
}

// DeleteOrder ...
func (o *OrderServiceClient) DeleteOrder(orderID string) *exception.Exception {
	_, err := o.client.DeleteOrder(context.Background(), &pb.DeleteOrderByOrderIDRequest{
		OrderID: orderID,
	})
	if err != nil {
		return &exception.Exception{
			Err:        err,
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
			StatusText: exception.STATUS_SERVICE_ERROR,
		}
	}

	return nil
}

// DeleteOrders ...
func (o *OrderServiceClient) DeleteOrders(userID string) *exception.Exception {
	_, err := o.client.DeleteOrders(context.Background(), &pb.DeleteOrdersByBuyerIDRequest{
		BuyerID: userID,
	})
	if err != nil {
		return &exception.Exception{
			Err:        err,
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
			StatusText: exception.STATUS_SERVICE_ERROR,
		}
	}

	return nil
}

func convertStringToTime(input string) *time.Time {
	parsedTime, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return nil
	}
	return &parsedTime
}
