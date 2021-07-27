package coingecko_service

import (
	coingecko_client "ejercicio-golang-meli-nueva/internal/client/coingecko"
	"ejercicio-golang-meli-nueva/internal/service"
)

type CoinGeckoService struct {
	CoinGeckoClient coingecko_client.CoinGeckoClient
}

func NewCoinGeckoService(coinGeckoClient *coingecko_client.CoinGeckoClient) (*CoinGeckoService) {
	return &CoinGeckoService{
		CoinGeckoClient: *coinGeckoClient,
	}
}

func (c CoinGeckoService) GetCurrentPrice(id string) (*service.CryptoResponse, error) {
	clientResponse, err := c.CoinGeckoClient.GetCoinPrice(id)
	if err != nil {
		return &service.CryptoResponse{
			Id:id,
			Partial: true,
		}, err
	}
	return &service.CryptoResponse{
		Id: id,
		Content: &service.Content{
			Price: clientResponse.MarketData.CurrentPrice["usd"],
			Currency: "usd",
		},
		Partial: false,
	}, nil
}