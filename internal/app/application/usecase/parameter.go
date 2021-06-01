package usecase

import (
	"github.com/airdb/scf-go/internal/app/domain"
	"github.com/airdb/scf-go/internal/app/domain/repository"
)

// Parameter is the usecase of getting parameter
func Parameter(r repository.IParameter) domain.Parameter {
	return r.Get()
}
