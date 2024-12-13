package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kreipikc/golang-gin-api/pkg/common/models"
)

type RegisterRequestBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h handler) RegisterUser(ctx *gin.Context) {
	body := RegisterRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	var user models.User

	user.Name = body.Name
	user.Email = body.Email
	user.Password = HashPass(body.Password)
	user.IsUser = true
	user.IsAdmin = false

	if result := h.DB.Create(&user); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &user)
}
