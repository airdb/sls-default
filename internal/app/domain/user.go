package domain

import "github.com/airdb/scf-go/internal/app/domain/valueobject"

// User is an order which has id, payment and person info
type User struct {
	ID      string
	Payment valueobject.Payment
}
