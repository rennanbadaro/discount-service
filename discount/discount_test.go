package discount_test

import (
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/rennanbadaro/discount-calculator/discount"
	"github.com/rennanbadaro/discount-calculator/infrastructure/repositories"
	"github.com/stretchr/testify/assert"
)

var fakeUser repositories.User
var fakeProduct repositories.Product

func setupBaseFakes() {
	baseUser := repositories.User{
		Id:          "fakeUserID",
		FirstName:   "John",
		LastName:    "Doe",
		DateOfBirth: time.Now().Add(time.Hour * 48).Format("2006-01-02"),
	}

	baseProduct := repositories.Product{
		Id:           "fakeProductID",
		PriceInCents: 1500,
		Title:        "Fake Title",
		Description:  "lorem lorem",
	}

	fakeUser = baseUser
	fakeProduct = baseProduct
}

func setNowToBlackfriday() {
	monkey.Patch(time.Now, func() time.Time {
		fakeTime, _ := time.Parse("2006-01-02", "2006-11-25")

		return fakeTime
	})
}

func resetNow() {
	monkey.Unpatch(time.Now)
}

type userRepoStub struct {
}

type productRepoStub struct {
}

func (u *userRepoStub) FetchByID(id string) (*repositories.User, error) {
	return &fakeUser, nil
}

func (u *productRepoStub) FetchByID(id string) (*repositories.Product, error) {
	return &fakeProduct, nil
}

func TestNoDiscount(t *testing.T) {
	setupBaseFakes()
	service := discount.NewDiscountService(&userRepoStub{}, &productRepoStub{})

	result, _ := service.GetDiscount(fakeUser.Id, fakeProduct.Id)

	expected := discount.Discount{
		Percentage:   0,
		ValueInCents: 0,
	}

	assert.Equal(t, &expected, result)
}

func TestBirthdayDiscount(t *testing.T) {
	setupBaseFakes()
	fakeUser.DateOfBirth = time.Now().Format("2006-01-02")

	service := discount.NewDiscountService(&userRepoStub{}, &productRepoStub{})

	result, _ := service.GetDiscount(fakeUser.Id, fakeProduct.Id)

	expected := discount.Discount{
		Percentage:   5,
		ValueInCents: int32(float32(fakeProduct.PriceInCents) * 0.05),
	}

	assert.Equal(t, &expected, result)
}

func TestBlackFridayDiscount(t *testing.T) {
	setNowToBlackfriday()
	setupBaseFakes()
	fakeUser.DateOfBirth = time.Now().Format("2006-01-02")

	service := discount.NewDiscountService(&userRepoStub{}, &productRepoStub{})

	result, _ := service.GetDiscount(fakeUser.Id, fakeProduct.Id)

	expected := discount.Discount{
		Percentage:   10,
		ValueInCents: int32(float32(fakeProduct.PriceInCents) * 0.1),
	}

	resetNow()

	assert.Equal(t, &expected, result)
}

func TestMaximumDiscount(t *testing.T) {
	setNowToBlackfriday()
	setupBaseFakes()
	fakeUser.DateOfBirth = time.Now().Format("2006-01-02")

	service := discount.NewDiscountService(&userRepoStub{}, &productRepoStub{})

	result, _ := service.GetDiscount(fakeUser.Id, fakeProduct.Id)

	expected := discount.Discount{
		Percentage:   10,
		ValueInCents: int32(float32(fakeProduct.PriceInCents) * 0.1),
	}

	assert.Equal(t, &expected, result)
}
