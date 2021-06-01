package repository

import (
	"errors"

	"github.com/airdb/scf-go/internal/app/adapter/mysql"
	"github.com/airdb/scf-go/internal/app/adapter/mysql/model"
	"github.com/airdb/scf-go/internal/app/domain"
	"github.com/airdb/scf-go/internal/app/domain/factory"
	"gorm.io/gorm"
)

// User is the repository of domain.User
type User struct{}

// Get gets order
func (u User) Get() domain.User {
	db := mysql.Connection()
	var order model.Order
	// User has Person/Payment relation and Payment has Card relation which has CardBrand relation.
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
func (u User) Update(order domain.User) {
	db := mysql.Connection()

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
