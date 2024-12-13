package users

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kreipikc/golang-gin-api/pkg/common/models"
)

func (h handler) RefreshToken(ctx *gin.Context) {
	claims := CheckToken(ctx, "refresh_token")
	if claims == nil {
		return
	}

	var user models.User

	if result := h.DB.Where("email = ?", claims.Email).First(&user); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	AccessToken, AccessToken_expirationTime, err := CreateAccessToken(user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("access_token", AccessToken, int(time.Until(AccessToken_expirationTime).Seconds()), "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{"access_token": AccessToken})
}
