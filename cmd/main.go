package main

import (
	coingecko_client "ejercicio-golang-meli-nueva/internal/client/coingecko"
	"ejercicio-golang-meli-nueva/internal/server"
	"ejercicio-golang-meli-nueva/internal/server/controller/crypto_controller"
	"ejercicio-golang-meli-nueva/internal/service/coingecko_service"
)

func main() {
	BaseUrl := "https://api.coingecko.com/api/v3/coins/"
	coingeckoClient := coingecko_client.NewCoinGeckoClient(BaseUrl)
	coingeckoService := coingecko_service.NewCoinGeckoService(coingeckoClient)
	cryptoController := crypto_controller.NewCryptoController(coingeckoService)
	router := server.GetMeliServer(cryptoController)
	router.Run(":8080")
}