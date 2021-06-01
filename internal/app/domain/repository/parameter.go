package repository

import "github.com/airdb/scf-go/internal/app/domain"

// IParameter is interface of parameter repository
type IParameter interface {
	Get() domain.Parameter
}
