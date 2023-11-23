package grpchandler

import (
	"oms/pkg/domain"

	pb "github.myproto.com"
)

// GrpcHandler ...
type GrpcHandler struct {
	pb.UnimplementedOrderManagementServer
	domain domain.Order
}

// NewGrpcHandler ...
func NewGrpcHandler(
	domain domain.Order,
) *GrpcHandler {
	return &GrpcHandler{
		domain: domain,
	}
}
