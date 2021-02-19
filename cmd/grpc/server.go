package grpc

import (
	"fmt"
	"net"
	"os"

	"github.com/rennanbadaro/discount-service/infrastructure/proto"
	"google.golang.org/grpc"
)

func StartServer() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))

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
