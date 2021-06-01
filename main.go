package main

import (
	"time"

	"github.com/airdb/scf-go/internal/app/adapter"
	"github.com/airdb/scf-go/internal/version"
)

func main() {
	// Init the loc.
	loc, _ := time.LoadLocation("Asia/Shanghai")

	// Set timezone.
	version.CreatedAt = time.Now().In(loc)

	adapter.NewRouter()
}
