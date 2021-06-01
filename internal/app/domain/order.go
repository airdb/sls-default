package domain

import "github.com/airdb/scf-go/internal/app/domain/valueobject"

// Order is an order which has id, payment and person info
type Order struct {
	ID      string
	Payment valueobject.Payment
	Person  Person
}
