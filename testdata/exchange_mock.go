package testdata

import "github.com/airdb/sls-default/internal/app/domain/valueobject"

// MExchange is mock of service.IExchange
type MExchange struct{}

// Ticker is mock implementation of service.IExchange.Ticker()
func (e MExchange) Ticker(p valueobject.Pair) valueobject.Ticker {
	return valueobject.Ticker{
		Sell:      "1000",
		Buy:       "1000",
		High:      "2000",
		Low:       "500",
		Last:      "1200",
		Vol:       "20",
		Timestamp: "1600769562",
	}
}

func (e MExchange) GetUser() valueobject.User {
	return valueobject.User{}
}
