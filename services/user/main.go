package main

import (
	"fmt"
	usergrpc "github.com/b4ckslash/rental-app/services/user/grpc"
	"github.com/b4ckslash/rental-app/services/user/delivery"
	"github.com/b4ckslash/rental-app/services/user/entity"
	"github.com/b4ckslash/rental-app/services/user/repository"
	"github.com/b4ckslash/rental-app/services/user/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost", "postgres", "123123123", "user_db", "5432",
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&entity.User{})

	repo := repository.NewUserRepository(db)
	uc := usecase.NewUserUsecase(repo)

	go usergrpc.StartUserGRPC(repo)

	r := gin.Default()
	delivery.RegisterRoutes(r, uc)
	r.Run(":8082")
}
