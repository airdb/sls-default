package main

import (
	"time"

	"github.com/airdb/scf-go/api"
	"github.com/airdb/scf-go/internal/version"
)

func main() {
	// Init the loc.
	loc, _ := time.LoadLocation("Asia/Shanghai")

	// Set timezone.
	version.CreatedAt = time.Now().In(loc)

	api.NewRouter()
}
