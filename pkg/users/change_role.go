package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kreipikc/golang-gin-api/pkg/common/models"
)

type ChangeRoleRequestBody struct {
	Role string `json:"role"`
}

func (h handler) ChangeRole(ctx *gin.Context) {
	claims := CheckToken(ctx, "access_token")
	if claims == nil {
		return
	}

	// Поиск пользователя в БД
	var user models.User
	if result := h.DB.Where("email = ?", claims.Email).First(&user); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Проверка на роль админа
	if !user.IsAdmin {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Not enough rights"})
		return
	}

	id := ctx.Param("id")
	body := ChangeRoleRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var new_role_user models.User

	if result := h.DB.First(&new_role_user, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	new_role_user.IsUser = false
	new_role_user.IsAdmin = false

	if body.Role == "admin" {
		new_role_user.IsAdmin = true
	} else if body.Role == "user" {
		new_role_user.IsUser = true
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Role is not valid"})
	}

	h.DB.Save(&new_role_user)

	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
}
