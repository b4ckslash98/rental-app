package main

import (
	"fmt"
	"github.com/b4ckslash/rental-app/services/order/delivery"
	"github.com/b4ckslash/rental-app/services/order/entity"
	"github.com/b4ckslash/rental-app/services/order/grpc/carclient"
	"github.com/b4ckslash/rental-app/services/order/grpc/userclient"
	"github.com/b4ckslash/rental-app/services/order/repository"
	"github.com/b4ckslash/rental-app/services/order/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost", "postgres", "123123123", "order_db", "5432",
	)
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&entity.Order{})

	repo := repository.NewOrderRepository(db)
	userClient := userclient.New()
	carClient := carclient.New()
	uc := usecase.NewOrderUsecase(repo, userClient, carClient)

	r := gin.Default()
	delivery.RegisterRoutes(r, uc)
	r.Run(":8083")
}
