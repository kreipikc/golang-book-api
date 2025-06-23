package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kreipikc/golang-gin-api/pkg/books"
	"github.com/kreipikc/golang-gin-api/pkg/common/database"
	"github.com/kreipikc/golang-gin-api/pkg/users"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := database.Init(dbUrl)

	books.RegisterRouter(r, h)
	users.RegisterRouter(r, h)

	r.Run(port)
}
