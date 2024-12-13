package users

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kreipikc/golang-gin-api/pkg/common/models"
	"github.com/spf13/viper"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var SECRET_KEY_JWT []byte

func InitJWT() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	secretKey := viper.GetString("SECRET_KEY_JWT")
	if secretKey == "" {
		panic("SECRET_KEY_JWT is not set in the configuration file")
	}
	SECRET_KEY_JWT = []byte(secretKey)
}

func (h handler) LoginUser(ctx *gin.Context) {
	body := LoginRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User

	if result := h.DB.Where("email = ?", body.Email).First(&user); result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if !VerifyPass(body.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	InitJWT()
	AccessToken, AccessToken_expirationTime, err := CreateAccessToken(user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	RefreshToken, RefreshToken_expirationTime, err := CreateRefreshToken(user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("access_token", AccessToken, int(time.Until(AccessToken_expirationTime).Seconds()), "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", RefreshToken, int(time.Until(RefreshToken_expirationTime).Seconds()), "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{"message": "Logged in successfully", "access_token": AccessToken, "refresh_token": RefreshToken})
}
