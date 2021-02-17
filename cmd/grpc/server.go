package grpc

import (
	"net"

	"github.com/rennanbadaro/discount-calculator/infrastructure/proto"
	"google.golang.org/grpc"
)

func StartServer() error {
	listener, err := net.Listen("tcp", ":9000")

	if err != nil {
		return err
	}

	server := grpc.NewServer()

	controller, err := NewDiscountController()

	if err != nil {
		return err
	}

	proto.RegisterDiscountServiceServer(server, controller)

	if err := server.Serve(listener); err != nil {
		return err
	}

	return nil
}
