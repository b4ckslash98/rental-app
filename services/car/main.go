package main

import (
	"fmt"
	"github.com/b4ckslash/rental-app/services/car/delivery"
	"github.com/b4ckslash/rental-app/services/car/entity"
	cargrpc "github.com/b4ckslash/rental-app/services/car/grpc"
	"github.com/b4ckslash/rental-app/services/car/repository"
	"github.com/b4ckslash/rental-app/services/car/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost", "postgres", "123123123", "car_db", "5432",
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entity.Car{})

	carRepo := repository.NewCarRepository(db)
	carUC := usecase.NewCarUsecase(carRepo)

	go cargrpc.StartCarGRPC(carRepo)

	r := gin.Default()
	delivery.RegisterRoutes(r, carUC)
	r.Run(":8081")
}
