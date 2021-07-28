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

func (cr *CryptoController) CoinPrices(c *gin.Context) {
	ids := []string{"bitcoin", "lovehearts1", "lucent"}
	response, err := cr.CryptoService.GetCurrentPrices(ids)
	if err != nil {
		c.JSON(http.StatusPartialContent, response)
		return
	}
	for _, v := range response {
		if (v.Partial) {
			c.JSON(http.StatusPartialContent, response)
			return
		}
	}
	c.IndentedJSON(http.StatusOK, response)
}