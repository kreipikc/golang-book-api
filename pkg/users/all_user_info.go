package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kreipikc/golang-gin-api/pkg/common/models"
)

func (h handler) GetAllInfo(ctx *gin.Context) {
	claims := CheckToken(ctx, "access_token")
	if claims == nil {
		return
	}

	var user models.User
	if result := h.DB.Where("email = ?", claims.Email).First(&user); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if !user.IsAdmin {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Not enough rights"})
		return
	}

	var users []models.User
	if result := h.DB.Find(&users); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &users)
}
