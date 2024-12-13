package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kreipikc/golang-gin-api/pkg/common/models"
)

func (h handler) GetInfoMe(ctx *gin.Context) {
	// tokenString, err := ctx.Cookie("access_token")
	// if err != nil {
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No token found"})
	// 	return
	// }

	// claims := &JwtBody{}

	// token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
	// 	return SECRET_KEY_JWT, nil
	// })

	// if err != nil {
	// 	if err == jwt.ErrSignatureInvalid {
	// 		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token parsing error"})
	// 	return
	// }

	// if !token.Valid {
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
	// 	return
	// }

	claims := CheckToken(ctx, "access_token")
	if claims == nil {
		return
	}

	var user models.User

	if result := h.DB.Where("email = ?", claims.Email).First(&user); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, &user)
}
