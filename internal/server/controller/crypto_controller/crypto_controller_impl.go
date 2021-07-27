package crypto_controller

import (
	"ejercicio-golang-meli-nueva/internal/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type CryptoController struct {
	CryptoService service.CryptoService
}

func NewCryptoController(service service.CryptoService) *CryptoController {
	return &CryptoController{
		CryptoService: service,
	}
}

func (cr *CryptoController) CoinPrice(c *gin.Context) {
	id := c.Query("id")
	response, err := cr.CryptoService.GetCurrentPrice(id)
	if err != nil {
		c.JSON(http.StatusPartialContent, response)
		return
	}
	c.JSON(http.StatusOK, response)
}