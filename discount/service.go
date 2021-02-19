package discount

import (
	"time"

	"github.com/rennanbadaro/discount-service/infrastructure/repositories"
)

type Discount struct {
	Percentage   float32
	ValueInCents int32
}

type IDiscountService interface {
	GetDiscount(userId, productId string) (*Discount, error)
}

type DiscountService struct {
	userRepository    repositories.IUserRepository
	productRepository repositories.IProductRepository
}

func (d *DiscountService) GetDiscount(userId, productId string) (*Discount, error) {
	product, err := d.productRepository.FetchByID(productId)

	if err != nil {
		return nil, err
	}

	isBlackfriday := time.Now().Month() == time.November && time.Now().Day() == 25

	if isBlackfriday {
		return &Discount{
			Percentage:   10,
			ValueInCents: int32(float32(product.PriceInCents) * 0.1),
		}, nil
	}

	user, err := d.userRepository.FetchByID(userId)

	if err != nil {
		return nil, err
	}

	isUsersBirthday := time.Now().Format("2006-01-02") == user.DateOfBirth

	if isUsersBirthday {
		return &Discount{
			Percentage:   5,
			ValueInCents: int32(float32(product.PriceInCents) * 0.05),
		}, nil
	}

	return &Discount{
		Percentage:   0,
		ValueInCents: 0,
	}, nil
}

func NewDiscountService(
	userRepo repositories.IUserRepository,
	productRepo repositories.IProductRepository,
) *DiscountService {
	return &DiscountService{
		userRepository:    userRepo,
		productRepository: productRepo,
	}
}
