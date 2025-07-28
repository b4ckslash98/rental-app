package usecase

import (
	"github.com/b4ckslash/rental-app/services/car/entity"
	"github.com/b4ckslash/rental-app/services/car/repository"
)

type CarUsecase interface {
	CreateCar(car *entity.Car) error
	GetCarByID(id int) (*entity.Car, error)
	UpdateCar(car *entity.Car) error
	DeleteCar(id int) error
	ListCars() ([]*entity.Car, error)
}

type carUsecase struct {
	repo repository.CarRepository
}

func NewCarUsecase(r repository.CarRepository) CarUsecase {
	return &carUsecase{repo: r}
}

func (uc *carUsecase) CreateCar(car *entity.Car) error {
	return uc.repo.Save(car)
}

func (uc *carUsecase) GetCarByID(id int) (*entity.Car, error) {
	return uc.repo.FindByID(id)
}

func (uc *carUsecase) UpdateCar(car *entity.Car) error {
	return uc.repo.Update(car)
}

func (uc *carUsecase) DeleteCar(id int) error {
	return uc.repo.Delete(id)
}

func (uc *carUsecase) ListCars() ([]*entity.Car, error) {
	return uc.repo.FindAll()
}
