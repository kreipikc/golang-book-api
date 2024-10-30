package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kreipikc/golang-gin-api/pkg/common/models"
)

func (h handler) GetBook(ctx *gin.Context) {
	id := ctx.Param("id")

	var book models.Book

	if result := h.DB.Where("id = ?", id).First(&book); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &book)
}
