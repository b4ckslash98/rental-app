package delivery

import (
	"net/http"

	"github.com/b4ckslash98/rental-app/services/user/usecase"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, uc usecase.UserUsecase) {
	user := r.Group("/users")

	user.POST("/register", func(c *gin.Context) {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
			Role     string `json:"role"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := uc.Register(req.Email, req.Password, req.Role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "registered"})
	})

	user.POST("/login", func(c *gin.Context) {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := uc.Login(req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}
		// Here should generate JWT token, skipped for now
		c.JSON(http.StatusOK, gin.H{"user": user.Email, "role": user.Role})
	})
}
