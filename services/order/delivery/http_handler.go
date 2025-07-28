package delivery

import (
	"net/http"

	"github.com/b4ckslash/rental-app/services/order/entity"
	"github.com/b4ckslash/rental-app/services/order/usecase"
	mw "github.com/b4ckslash/rental-app/services/user/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, uc usecase.OrderUsecase) {
	order := r.Group("/orders")
	order.Use(mw.JWTAuthMiddleware())

	order.POST("", mw.CustomerOnly(), func(c *gin.Context) {
		var input entity.Order
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		uid := c.GetInt("user_id")
		input.UserID = uid

		if err := uc.BookOrder(input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "order placed"})
	})

	order.GET("", func(c *gin.Context) {
		uid := c.GetInt("user_id")
		role := c.GetString("role")
		result, err := uc.ListOrders(uid, role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	})
}
