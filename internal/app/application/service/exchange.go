package service

import "github.com/airdb/sls-default/internal/app/domain/valueobject"

// IExchange is interface of bitcoin exchange
type IExchange interface {
	GetUser() valueobject.User
	Ticker(p valueobject.Pair) valueobject.Ticker
}
