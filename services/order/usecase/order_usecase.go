package usecase

import (
	"errors"
	"time"

	"github.com/b4ckslash/rental-app/services/order/entity"
	"github.com/b4ckslash/rental-app/services/order/grpc/carclient"
	"github.com/b4ckslash/rental-app/services/order/grpc/userclient"
	"github.com/b4ckslash/rental-app/services/order/repository"
)

type OrderUsecase interface {
	BookOrder(input entity.Order) error
	ListOrders(userID int, role string) ([]*entity.Order, error)
}

type orderUsecase struct {
	repo       repository.OrderRepository
	userClient userclient.UserGRPCClient
	carClient  carclient.CarGRPCClient
}

func NewOrderUsecase(r repository.OrderRepository, uc userclient.UserGRPCClient, cc carclient.CarGRPCClient) OrderUsecase {
	return &orderUsecase{r, uc, cc}
}

func (uc *orderUsecase) BookOrder(order entity.Order) error {
	errs := make(chan error, 2)

	go func() {
		if ok := uc.userClient.ValidateUser(order.UserID); !ok {
			errs <- errors.New("invalid user")
		} else {
			errs <- nil
		}
	}()

	go func() {
		if ok := uc.carClient.CheckAvailability(order.CarID); !ok {
			errs <- errors.New("car not available")
		} else {
			errs <- nil
		}
	}()

	for i := 0; i < 2; i++ {
		if err := <-errs; err != nil {
			return err // cancel saga
		}
	}

	order.OrderDate = time.Now().Format("2006-01-02")
	if err := uc.repo.Save(&order); err != nil {
		return err
	}

	go func() {
		// simulate notification
		time.Sleep(1 * time.Second)
		println("Notification sent to user", order.UserID)
	}()

	return nil
}

func (uc *orderUsecase) ListOrders(userID int, role string) ([]*entity.Order, error) {
	if role == "admin" {
		return uc.repo.FindAll()
	}
	return uc.repo.FindByUser(userID)
}
