package service

type CryptoResponse struct {
	Id        	string      `json:"id"`
	Content		*Content	`json:"content,omitempty"`
	Partial 	bool		`json:"partial"`
}

type Content struct {
	Price		float64		`json:"price"`
	Currency	string		`json:"currency"`
}

type CryptoService interface {
	GetCurrentPrice(id string) (*CryptoResponse, error)
}