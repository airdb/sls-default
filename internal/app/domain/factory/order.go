package factory

import (
	"github.com/airdb/scf-go/internal/app/domain"
)

// Order is the factory of domain.Order
type Order struct{}

// Generate generates domain.Order from primitives
func (of Order) Generate(
	orderID string,
) domain.Order {
	return domain.Order{
		ID: orderID,
	}
}
