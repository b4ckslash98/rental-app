package repository

import (
    "gorm.io/gorm"
    "github.com/b4ckslash/rental-app/services/car/entity"
)

type CarRepository interface {
    Save(car *entity.Car) error
    FindByID(id int) (*entity.Car, error)
    Update(car *entity.Car) error
    Delete(id int) error
    FindAll() ([]*entity.Car, error)
}

type carRepo struct {
    db *gorm.DB
}

func NewCarRepository(db *gorm.DB) CarRepository {
    return &carRepo{db}
}

func (r *carRepo) Save(car *entity.Car) error {
    return r.db.Create(car).Error
}

func (r *carRepo) FindByID(id int) (*entity.Car, error) {
    var car entity.Car
    err := r.db.First(&car, id).Error
    return &car, err
}

func (r *carRepo) Update(car *entity.Car) error {
    return r.db.Save(car).Error
}

func (r *carRepo) Delete(id int) error {
    return r.db.Delete(&entity.Car{}, id).Error
}

func (r *carRepo) FindAll() ([]*entity.Car, error) {
    var cars []*entity.Car
    err := r.db.Find(&cars).Error
    return cars, err
}