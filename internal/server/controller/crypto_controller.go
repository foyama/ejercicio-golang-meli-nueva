package controller

import (
	"github.com/gin-gonic/gin"
)

type CryptoController interface {
	CoinPrice(c *gin.Context)
}