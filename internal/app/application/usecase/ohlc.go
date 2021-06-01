package usecase

import (
	"github.com/airdb/scf-go/internal/app/application/service"
	"github.com/airdb/scf-go/internal/app/domain/valueobject"
)

// OhlcArgs are arguments of Ohlc usecase
type OhlcArgs struct {
	E service.IExchange
	P valueobject.Pair
	T valueobject.Timeunit
}
