package users

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type JwtBody struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func HashPass(pass string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password:", err)
	}
	return string(hashedPassword)
}

func VerifyPass(pass string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(pass))
	if err != nil {
		return false
	} else {
		return true
	}
}

func CreateAccessToken(email string) (string, time.Time, error) {
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &JwtBody{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SECRET_KEY_JWT)
	return tokenString, expirationTime, err
}

func CreateRefreshToken(email string) (string, time.Time, error) {
	expirationTime := time.Now().Add(30 * 24 * time.Hour)
	claims := &JwtBody{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SECRET_KEY_JWT)
	return tokenString, expirationTime, err
}

func CheckToken(ctx *gin.Context, typeToken string) *JwtBody {
	tokenString, err := ctx.Cookie(typeToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No token found"})
		return nil
	}

	claims := &JwtBody{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return SECRET_KEY_JWT, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return nil
		}
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token parsing error"})
		return nil
	}

	if !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return nil
	}
	return claims
}
