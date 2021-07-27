package server

import (
	"ejercicio-golang-meli-nueva/internal/server/controller"

	"github.com/gin-gonic/gin"
)

func GetMeliServer(input controller.CryptoController) *gin.Engine{
	router := gin.Default()
	m := router.Group("/meli")
	m.GET("/coinprice", input.CoinPrice)
	m.GET("/coinprices", input.CoinPrices)
	return router
}