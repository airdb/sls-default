package repository

import "github.com/airdb/scf-go/internal/app/domain"

// IOrder is interface of order repository
type IOrder interface {
	Get() domain.User
	Update(domain.User)
}
