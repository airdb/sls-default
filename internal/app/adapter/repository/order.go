package repository

import (
	"errors"

	"github.com/airdb/scf-go/internal/app/adapter/postgresql"
	"github.com/airdb/scf-go/internal/app/adapter/postgresql/model"
	"github.com/airdb/scf-go/internal/app/domain"
	"github.com/airdb/scf-go/internal/app/domain/factory"
	"gorm.io/gorm"
)

// Order is the repository of domain.Order
type Order struct{}

// Get gets order
func (o Order) Get() domain.Order {
	db := postgresql.Connection()
	var order model.Order
	// Order has Person/Payment relation and Payment has Card relation which has CardBrand relation.
	result := db.Preload("Person").Preload("Payment.Card.CardBrand").Find(&order)
	if result.Error != nil {
		panic(result.Error)
	}
	orderFactory := factory.Order{}
	return orderFactory.Generate(
		order.ID,
	)
}

// Update updates order
func (o Order) Update(order domain.Order) {
	db := postgresql.Connection()

	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Exec("update payments set card_id = 1").Error
		if err != nil {
			return errors.New("rollback")
		}
		return nil // commit
	})
	if err != nil {
		panic(err)
	}
}
