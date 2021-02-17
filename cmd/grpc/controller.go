package grpc

import (
	"context"
	"log"

	"github.com/rennanbadaro/discount-calculator/discount"
	"github.com/rennanbadaro/discount-calculator/infrastructure/proto"
	"github.com/rennanbadaro/discount-calculator/infrastructure/repositories"
	"github.com/rennanbadaro/discount-calculator/infrastructure/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DiscountController struct {
	proto.UnimplementedDiscountServiceServer

	service *discount.DiscountService
}

func (dc *DiscountController) GetDiscount(
	ctx context.Context,
	in *proto.GetDiscountRequest,
) (*proto.GetDiscountResponse, error) {
	discount, err := dc.service.GetDiscount(in.UserId, in.ProductId)

	if err != nil {
		log.Println(err)
		return &proto.GetDiscountResponse{}, status.Error(codes.Internal, "error getting discount")
	}

	response := proto.Discount{Percentage: discount.Percentage, ValueInCents: discount.ValueInCents}

	return &proto.GetDiscountResponse{Discount: &response}, nil
}

func NewDiscountController() (*DiscountController, error) {
	pgClient, err := storage.NewPostgresClient()

	if err != nil {
		return nil, err
	}

	userRepo := repositories.NewUserRepository(pgClient)
	productRepo := repositories.NewProductRepository(pgClient)
	service := discount.NewDiscountService(userRepo, productRepo)

	return &DiscountController{service: service}, nil
}
