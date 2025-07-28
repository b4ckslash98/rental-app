package delivery

import (
	"net/http"
	"strconv"

	"github.com/b4ckslash98/rental-app/services/car/entity"
	"github.com/b4ckslash98/rental-app/services/car/usecase"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, carUC usecase.CarUsecase) {
	cars := r.Group("/cars")

	cars.POST("", func(c *gin.Context) {
		var car entity.Car
		if err := c.ShouldBindJSON(&car); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := carUC.CreateCar(&car)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, car)
	})

	cars.GET("", func(c *gin.Context) {
		result, err := carUC.ListCars()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	})

	cars.GET(":id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		car, err := carUC.GetCarByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
			return
		}
		c.JSON(http.StatusOK, car)
	})

	cars.PUT(":id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var car entity.Car
		if err := c.ShouldBindJSON(&car); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		car.ID = id
		err := carUC.UpdateCar(&car)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, car)
	})

	cars.DELETE(":id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		err := carUC.DeleteCar(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
	})
}
