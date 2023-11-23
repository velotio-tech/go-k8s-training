package grpc

import (
	"log"
	"net"

	pb "github.myproto.com"

	"google.golang.org/grpc"
)

// Server ...
type Server struct {
	port       string
	grpcServer pb.OrderManagementServer
}

// NewGrpcServer ...
func NewGrpcServer(port string, grpcServer pb.OrderManagementServer) *Server {
	return &Server{port: port, grpcServer: grpcServer}
}

func (s *Server) ListenAndServe() error {
	lis, err := net.Listen("tcp", s.port)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterOrderManagementServer(grpcServer, s.grpcServer)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Println("failed to serve grpc server : ", err)
		return err
	}

	return nil
}
