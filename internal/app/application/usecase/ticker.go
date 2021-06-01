package usecase

import (
	"github.com/airdb/scf-go/internal/app/application/service"
	"github.com/airdb/scf-go/internal/app/domain/valueobject"
)

// Ticker is the usecase of getting ticker
func Ticker(e service.IExchange, p valueobject.Pair) valueobject.Ticker {
	return e.Ticker(p)
}

func GetUser(e service.IExchange) valueobject.User {

	return e.GetUser()
}
