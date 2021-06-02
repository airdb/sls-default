package service

import (
	"github.com/airdb/scf-go/internal/app/domain/valueobject"
)

// IExchange is interface of bitcoin exchange
type IExchange interface {
	GetUser() valueobject.User
	Ticker(p valueobject.Pair) valueobject.Ticker
}

func (b Bitbank) GetUser() valueobject.User {
	return valueobject.User{
		ID:       1,
		Username: "dean",
	}
}

// Bitbank is an bitcoin exchange
type Bitbank struct{}

// Ticker gets ticker via bitbank public api
func (b Bitbank) Ticker(p valueobject.Pair) valueobject.Ticker {
	return valueobject.Ticker{
		Timestamp: "2016",
	}
}

/*
func convertToCandlestick(res bitbankohlcresponse) []valueobject.CandleStick {
	ohlcs := res.Data.Candlestick[0].Ohlcv
	cs := make([]valueobject.CandleStick, 0)
	for _, v := range ohlcs {
		var base1000 float64 = 1000
		base10 := 10
		timestamp := strconv.FormatInt(int64(v[5].(float64)/base1000), base10)
		cs = append(cs, valueobject.CandleStick{
			Open:      v[0].(string),
			High:      v[1].(string),
			Low:       v[2].(string),
			Close:     v[3].(string),
			Volume:    v[4].(string),
			Timestamp: timestamp,
		})
	}
	return cs
}
*/

// Ticker is the usecase of getting ticker
func Ticker(e IExchange, p valueobject.Pair) valueobject.Ticker {
	return e.Ticker(p)
}

func GetUser(e IExchange) valueobject.User {
	return e.GetUser()
}
