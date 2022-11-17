package order

import (
	"context"

	"github.com/velotio-ajaykumbhar/microservice/order/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Database        = "microservice"
	OrderCollection = "order"
)

type OrderService interface {
	AddItem(context.Context, AddItemDTO) (string, error)
	GetAllOrder(context.Context, OrderAllDTO) ([]Order, error)
	GetOrder(context.Context, OrderDTO) (Order, error)
	DeleteAllOrder(context.Context, OrderAllDTO) (string, error)
	DeleteOrder(context.Context, OrderDTO) (string, error)
}

type orderService struct {
	mongo *mongo.Client
}

func NewOrderService(ds *database.DataSource) OrderService {
	return &orderService{
		mongo: ds.Mongo,
	}
}

func (os orderService) AddItem(ctx context.Context, item AddItemDTO) (string, error) {
	orderCollection := os.mongo.Database(Database).Collection(OrderCollection)

	filter := bson.M{
		"orderId": item.OrderId,
		"userId":  item.UserId,
	}

	update := bson.M{
		"$push": bson.M{
			"items": item.Item,
		},
	}
	opts := options.Update().SetUpsert(true)

	_, err := orderCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return "", err
	}

	return "item add to order", nil
}

func (os orderService) GetAllOrder(ctx context.Context, orderAllDTO OrderAllDTO) ([]Order, error) {
	orderCollection := os.mongo.Database(Database).Collection(OrderCollection)

	filter := bson.M{
		"userId": orderAllDTO.UserId,
	}

	cursor, err := orderCollection.Find(ctx, filter)
	if err != nil {
		return []Order{}, err
	}

	orders := []Order{}
	err = cursor.All(ctx, &orders)
	if err != nil {
		return []Order{}, nil
	}

	return orders, nil
}

func (os orderService) GetOrder(ctx context.Context, orderDTO OrderDTO) (Order, error) {
	orderCollection := os.mongo.Database(Database).Collection(OrderCollection)

	filter := bson.M{
		"userId":  orderDTO.UserId,
		"orderId": orderDTO.OrderId,
	}

	order := Order{}
	err := orderCollection.FindOne(ctx, filter).Decode(&order)
	if err != nil {
		return Order{}, err
	}

	return order, nil
}

func (os orderService) DeleteAllOrder(ctx context.Context, orderAllDTO OrderAllDTO) (string, error) {
	orderCollection := os.mongo.Database(Database).Collection(OrderCollection)

	filter := bson.M{
		"userId": orderAllDTO.UserId,
	}

	_, err := orderCollection.DeleteMany(ctx, filter)
	if err != nil {
		return "", err
	}

	return "all record deleted sucessfully", nil
}

func (os orderService) DeleteOrder(ctx context.Context, orderDTO OrderDTO) (string, error) {
	orderCollection := os.mongo.Database(Database).Collection(OrderCollection)

	filter := bson.M{
		"userId":  orderDTO.UserId,
		"orderId": orderDTO.OrderId,
	}

	_, err := orderCollection.DeleteOne(ctx, filter)
	if err != nil {
		return "", err
	}

	return "record deleted sucessfully", nil
}
