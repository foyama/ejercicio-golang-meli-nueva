package coingecko_client

import (
	melihttp "ejercicio-golang-meli-nueva/pkg/http"
	//"net/http"
)

type CoinGeckoResponse struct {
	Id        	string       `json:"id"`
	MarketData 	CurrentPrice `json:"market_data"`
}

type CurrentPrice struct {
	CurrentPrice	map[string]float64  `json:"current_price"`
}

func NewCoinGeckResponse() *CoinGeckoResponse {
	return &CoinGeckoResponse{}
}

type CoinGeckoClient struct {
	Url string
}

func NewCoinGeckoClient(url string) *CoinGeckoClient {
	return &CoinGeckoClient{
		Url: url,
	}
}

func (c *CoinGeckoClient) GetCoinPrice(path string) (*CoinGeckoResponse, error) {
	endpoint := c.Url + path
	response, err := melihttp.DoGet(endpoint)
	if err != nil {
		return nil, err
	}

	return response, nil
}