package coingecko_client

import (
	"ejercicio-golang-meli-nueva/pkg/http"
	"encoding/json"
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
	response, err := http.DoGet(endpoint)
	if err != nil {
		return nil, err
	}
	coinGeckoResponse := NewCoinGeckResponse()
	err = json.Unmarshal(response.Body, &coinGeckoResponse)
	if err != nil {
		return nil, err
	}
	return coinGeckoResponse, nil
}