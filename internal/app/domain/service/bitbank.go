package service

import (
	"github.com/airdb/scf-go/internal/app/domain/valueobject"
)

func (b Bitbank)  GetUser() valueobject.User {
	return valueobject.User{
		ID:       111,
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

// Ohlc gets open, high, low, and close via bitbank public api
// NOTICE: This works from 0AM (UTC) due to its api constraints
func (b Bitbank) Ohlc(p valueobject.Pair, t valueobject.Timeunit) []valueobject.CandleStick {
	var resp []valueobject.CandleStick
	resp = append(resp, valueobject.CandleStick{})

	return resp
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

// OneMin is a timeunit
const (
	OneMin valueobject.Timeunit = iota
	FiveMin
	FifteenMin
	ThirtyMin
	OneHour
	FourHour
	EightHour
	TweleveHour
	OneDay
	OneWeek
)

func (b Bitbank) convertTimeunit(t valueobject.Timeunit) string {
	return b.timeunitNames()[t]
}

func (b Bitbank) timeunitNames() []string {
	return []string{
		"1min",
		"5min",
		"15min",
		"30min",
		"1hour",
		"4hour",
		"8hour",
		"12hour",
		"1day",
		"1week",
	}
}

func (b Bitbank) convertPair(p valueobject.Pair) string {
	return b.pairNames()[p]
}

func (b Bitbank) pairNames() []string {
	return []string{
		"btc_jpy",
	}
}

func (b Bitbank) yyyy(t valueobject.Timeunit) string {
	day := time.Now()
	var layout string
	switch t {
	case OneMin, FiveMin, FifteenMin, ThirtyMin, OneHour:
		layout = "20060102"
	default:
		layout = "2006"
	}
	return day.Format(layout)
}
*/
