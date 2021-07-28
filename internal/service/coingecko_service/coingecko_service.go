package coingecko_service

import (
	coingecko_client "ejercicio-golang-meli-nueva/internal/client/coingecko"
	"ejercicio-golang-meli-nueva/internal/service"
	"sync"
)

type CoinGeckoService struct {
	CoinGeckoClient coingecko_client.CoinGeckoClient
}

func NewCoinGeckoService(coinGeckoClient *coingecko_client.CoinGeckoClient) (*CoinGeckoService) {
	return &CoinGeckoService{
		CoinGeckoClient: *coinGeckoClient,
	}
}

func (c *CoinGeckoService) GetCurrentPrice(id string) (*service.CryptoResponse, error) {
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

func (c *CoinGeckoService) GetCurrentPrices(ids []string) ([]service.CryptoResponse, error) {
	concurrency := len(ids)
	channel := make(chan service.CryptoResponse, concurrency)
	var cryptoResponses []service.CryptoResponse
	var wg sync.WaitGroup
	wg.Add(concurrency)
	
	for _, id := range ids {
		go c.worker(id, channel, &wg)
	}
	wg.Wait()
	close(channel)

	for c := range channel {
		cryptoResponses = append(cryptoResponses, c)
	}
	return cryptoResponses, nil
}

func (s *CoinGeckoService) worker(id string, c chan <- service.CryptoResponse, wg *sync.WaitGroup) {
	defer wg.Done()
	defer s.Recover(id, c)
	response, err := s.CoinGeckoClient.GetCoinPrice(id)
	if err != nil {
		return
	}
	if (response.MarketData.CurrentPrice["usd"] == 0) {
		panic("No coin matched id:" + id)
	}
	c <- service.CryptoResponse{
		Id: id,
		Content: &service.Content{
			Price: response.MarketData.CurrentPrice["usd"],
			Currency: "usd",
		},
		Partial: false,
	}
}

func (s *CoinGeckoService) Recover(id string, c chan <- service.CryptoResponse) {
	if r := recover(); r != nil {
		c <- service.CryptoResponse{
			Id: id,
			Partial: true,
		}
	}
}