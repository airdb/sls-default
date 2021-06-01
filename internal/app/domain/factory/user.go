package factory

import (
	"github.com/airdb/scf-go/internal/app/domain"
)

// Order is the factory of domain.User
type Order struct{}

// Generate generates domain.User from primitives
func (of Order) Generate(
	orderID string,
) domain.User {
	return domain.User{
		ID: orderID,
	}
}
