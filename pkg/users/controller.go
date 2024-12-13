package users

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRouter(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/user")
	routes.GET("/me", h.GetInfoMe)
	routes.POST("/auth/register", h.RegisterUser)
	routes.POST("/auth/login", h.LoginUser)
	routes.POST("/auth/refresh", h.RefreshToken)
	routes.GET("/admin/get_all_info", h.GetAllInfo)
	routes.POST("/admin/change_role/:id", h.ChangeRole)
}
