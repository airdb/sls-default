package service

import "github.com/airdb/scf-go/internal/app/domain/valueobject"

// IExchange is interface of bitcoin exchange
type IExchange interface {
	GetUser() valueobject.User
	Ticker(p valueobject.Pair) valueobject.Ticker
}
